package search
import (
		"strings"
		// "fmt"
		"github.com/google/cayley/graph"
		// "net/url"
		"github.com/google/cayley"		 
		"../extractArticleData"
		 _ "github.com/google/cayley/graph/bolt"
		 "sort"
		"html/template"
		"bytes"
	)

type Article struct{
	Title string
	Description string
	Link string
}

const dbPath = "/Users/guptap/Documents/SearchTool/Database/backend.db"


func GetSearchResultsAsCards(data string,store *cayley.Handle)(string){
	tags:=strings.Split(data,",")
	pairs:=GetArticles(tags,store)
	if len(pairs)==0 {
		return NO_RESULTS
	}
	htmlC:=GetHtml(store,pairs)
	return htmlC
}



func GetHtml(store *cayley.Handle,pairs PairList)(string){
	
	var cards bytes.Buffer
	Articles:=make([]Article,0,cap(pairs))
	for _,pair:= range pairs{
		Articles=append(Articles,ConvertLinkToInfo(store,pair.Key))
	}
	t, err := template.New("foo").Parse(CARD_HTML)
	if err != nil {
		panic(err)
	}

	//
	type Wrapper struct {
		Articles []Article
	}

	err = t.Execute(&cards,Wrapper{Articles})
	if err != nil {
		panic(err)
	}
	return  cards.String() 
}



func GetArticles(tags []string,store *cayley.Handle)(PairList){
	links := make(map[string]int)

	for _,tag:=range tags{
		
		entities:=extractArticleData.GetEntityInfo(tag,"")

		if len(entities) == 0{
			continue
		}

		entitityInfo:=entities[0] // Get the First Entity //TODO:- Change to Best Label Match

			for _,category:=range entitityInfo.Categories{

				path := cayley.StartPath(store, category).
					In("has_category").
					In("has_entity")

				it := path.BuildIterator()
				it, _ = it.Optimize()

				 for cayley.RawNext(it) {
				 	//fmt.Println(category)
		   			link:=store.NameOf(it.Result())
					links[link]+=10
				  }
			}
	}
	return sortMapByValue(links)
}


type Node struct{
	Value int
	link string
	Right *Node
	Left *Node
}


func ConvertLinkToInfo(store *cayley.Handle,link string )(Article){
	path := cayley.StartPath(store, link).
		Save("has_title","title").
		Save("has_description","description")

	it := path.BuildIterator()
	it, _ = it.Optimize()
	
	graph.Next(it) 
	tags := make(map[string]graph.Value)
	it.TagResults(tags)
	return Article{
		store.NameOf(tags["title"]),
		store.NameOf(tags["description"]),
		link,
	}
}

type Pair struct {
  Key string
  Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

// A function to turn a map into a PairList, then sort and return it. 
func sortMapByValue(m map[string]int) PairList {
   p := make(PairList, len(m))
   i := 0
   for k, v := range m {
      p[i] = Pair{k, v}
      i+=1
   }
   sort.Sort(sort.Reverse(p))
   return p
}

// func main(){
// 	fmt.Println("asdg")
// }

const CARD_HTML=`
			<div class="row">
			{{with .Articles}}
				{{range .}}
				<div class="col s12 m6 l4">
                  <div class="card blue-grey darken-3 hoverable">
                    <div class="card-content white-text">
                      <span class="card-title">{{.Title}}</span>
                      <p>{{.Description}}</p>
                    </div>
                    <div class="card-action">
                      <a href={{.Link}}> Link </a>
                    </div>
                  </div>
                </div>
               {{end}}
              {{end}}
             </div>
               	`

const NO_RESULTS=`
<div class="row">
				<div class="col s12 m12 l12">
                  <div class="card red darken-1 hoverable">
                    <div class="card-content white-text">
                      <span class="card-title">Dang it!</span>
                      <p>No results found for this query.Try something else</p>
                    </div>
                    <div class="card-action">
                    </div>
                  </div>
                </div>
             </div>

`