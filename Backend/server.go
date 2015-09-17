package main

import (
	"./extractArticleData"
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	// "strconv"
	"github.com/google/cayley/graph"
	"github.com/google/cayley"
	 _ "github.com/google/cayley/graph/bolt"
	 "./search"


)

const cayleyHost ="localhost:64210"
const dbPath = "/Users/guptap/Documents/SearchTool/Database/backend.db"

type Article struct {
		Title string
		Description string
		Link string
		Image string
	}

/*
Calls the extractArticleData package to obtain Article Data and then adds the information to the database
*/
func submitLink(w http.ResponseWriter, r *http.Request,store *cayley.Handle) {
	
	data :=extractArticleData.GetDataFromArticle(r.FormValue("UploadLink"))

	store.AddQuad(cayley.Quad(data.Link, "has_parent", data.Parent, ""))
	store.AddQuad(cayley.Quad(data.Link, "has_date",  data.Date, ""))
	store.AddQuad(cayley.Quad(data.Link,"has_content",data.Content,""))
	store.AddQuad(cayley.Quad(data.Link,"has_title",data.Title,""))
	store.AddQuad(cayley.Quad(data.Link,"has_description",data.Description,""))
	store.AddQuad(cayley.Quad(data.Link,"has_image",data.Image,""))
	
	for _,entity := range data.Entities {

		store.AddQuad(cayley.Quad(data.Link,"has_entity",entity.Name,""))
		for _,class := range entity.Classes{
			store.AddQuad(cayley.Quad(entity.Name, "has_class",class,""))
			
		}

		for _,category := range entity.Categories{
			store.AddQuad(cayley.Quad(entity.Name, "has_category",category,""))
		}

	}


}

/*
	Display the Search Page
*/
func searchHandler(w http.ResponseWriter, r *http.Request,store *cayley.Handle){

	t, error := template.ParseFiles("website/Search/index.html")

	if error != nil {
		fmt.Println("error:", error)
		fmt.Println("HTML TEMPLATE WAS NOT FOUND. Check if the html file is present.")
	}

	t.Execute(w, nil)
}
/*
* Cals the Search package to search the database for the tags and returns Html with search results
*/
func searchResultsHandler(w http.ResponseWriter, r *http.Request,store *cayley.Handle){

	fmt.Fprintf(w,
		search.GetSearchResultsAsCards(r.URL.Query().Get("tags"),
			store) )
}


/*
	The root page.Gets Data from the database to Display articles
*/
func mainHandler(w http.ResponseWriter, r *http.Request,store *cayley.Handle){

	t, error := template.ParseFiles("website/templates/index.html")

	if error != nil {
		fmt.Println("error:", error)
		fmt.Println("HTML TEMPLATE WAS NOT FOUND. Check if the html file is present.")

	}
	
	path := cayley.StartPath(store, "Article").
		In().
		Tag("link").
		Save("has_image", "image").
		Save("has_title","title").
		Save("has_description","description")

	it := path.BuildIterator()
	it, _ = it.Optimize()
	
	articleList:= make([]Article, 0)

	for graph.Next(it) {
		tags := make(map[string]graph.Value)
		it.TagResults(tags)

		articleList=append(articleList,Article{
			store.NameOf(tags["title"]),
			store.NameOf(tags["description"]),
			store.NameOf(tags["link"]),
			store.NameOf(tags["image"]),
		})

	}

	type PageContent struct {
		Articles []Article	
	}

	t.Execute(w, PageContent{articleList})
}



func main() {

	store, err := cayley.NewGraph("bolt", dbPath, nil)
	if err != nil {
		fmt.Println("error in opening database", err)
	}

	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("website/images"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("website/fonts"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("website/assets"))))
	http.Handle("/searchAssets/", http.StripPrefix("/searchAssets/", http.FileServer(http.Dir("website/Search/assets"))))
	http.Handle("/sass/", http.StripPrefix("/sass/", http.FileServer(http.Dir("website/sass"))))
	
	http.HandleFunc("/submitPost", func(w http.ResponseWriter, r *http.Request) { submitLink(w, r, store) })
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) { mainHandler(w, r, store) })
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) { searchHandler(w, r, store) })
	http.HandleFunc("/searchResults", func(w http.ResponseWriter, r *http.Request) { searchResultsHandler(w, r, store) })
	fmt.Println("Server started at port 8080")
	http.ListenAndServe(":8080", nil)

}

/*
	Make's a POST http request
	@param Url=the url to which the call has to be made
	@param requestBody= the query or content body (json,etc) which you want to request.Needs to be in byte[] format.
	@return :- status String, []Byte of the result
*/
func httpRequest(url string, requestBody []byte) (string, []byte) {

	fmt.Println("URL:>", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return resp.Status, body

}
