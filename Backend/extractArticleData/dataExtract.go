package extractArticleData

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"net/url"
	"strings"
	"encoding/json"
)

const cherryAddress = "http://localhost:8081"
const dbpediaServerAddress = "http://localhost:1111/api/search/KeywordSearch?QueryString="

/*
* All the data that we extract from an article
 */
type ArticleData struct {
	Link        string //The web address of the article
	Date        string //The date the article was added to the daabase
	Parent      string // Will always be "ARTICLE"
	Content     string // The text content of the article
	Title       string //The Title of the web article
	Description string // A small descrption of the article
	Image       string // An Image If preset in the article
	Entities	[]EntityData //Entities present in the article
}
/*
*To get Dbpedia or freebaase Links
*/
type Disambiguatations struct{
	Dbpedia string
	Freebase string
}

// type Entity struct{
// 	Name string 
// 	Classes []string
// 	Categories []string
// }


/*
*Entity Structure Returned by Alchemy Api
*/


type EntityData struct{
	Name string `xml:"Label"`
	Classes []string `xml:"Classes>Class>Label"`
	Categories []string `xml:"Categories>Category>Label"`
}


type EntityDataList struct{
	 EntityDataSlice []EntityData `xml:"Result"`
}
/*
Get Article info(title,image,content,description),entities
*/
func GetDataFromArticle(linkUrl string) ArticleData {
	
	var allData ArticleData
	ExtractArticleTextInfo(&allData, &linkUrl)
	GetEntities(&allData)
	
	//fmt.Println(allData.Entities)
	allData.Parent="Article"
	allData.Link=linkUrl
	allData.Date=time.Now().Local().Format("2006-01-02 15:04:05 +0800")
	return allData
	
}

/**
* Extracts article information from the cherryPy server(which uses Goose) .
  @param allData *ArticleData : pointer to an articleData to where all information needs to be stored
*/
func ExtractArticleTextInfo(allData *ArticleData, linkUrl *string) {
	values:=url.Values{"url": {*linkUrl}}
	_, data := GetRequest(cherryAddress,"",values)
	//fmt.Println( string(data))
	err := json.Unmarshal(data, &allData)
	if err != nil {
		fmt.Println("Error in Extract Article Info json:", err)
	}
	//fmt.Println(allData)

}

/*
	Launch a Get Request. 
	@param baseUrl:= The base url
	@param path:= a path to add on to the base url. E.g=/test/api
	@param parameters:= parameters that you want to pass. 
						Pass an empty url.Values if no parameters are needed
*/
func GetRequest(baseUrl string,path string,urlValues url.Values) (string, []byte) {

	
	Url, err := url.Parse(baseUrl)
	if err != nil {
        panic("invalid Url for get request")
    }

	Url.Path += path
	
    Url.RawQuery = urlValues.Encode()
	
	fmt.Println("Get request to URL:>", Url.String())
	req, err := http.NewRequest("GET", Url.String(), nil)
	// req.Header.Set("X-Custom-Header", "myvalue")
	// req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return resp.Status, body
}

type Resource struct{
			//SurfaceForm string `xml:"surfaceForm,attr"`
			URI string `xml:"URI,attr"`
}
/*
* Get Entities from the DBPedia Spotlight server
*/

func GetEntities(allData *ArticleData){
		
		urlValues := url.Values{}
		urlValues.Set("text", allData.Content)
		urlValues.Add("powered_by", "no")
		urlValues.Add("policy", "whitelist")
		urlValues.Add("confidence", "0.5")
		status,data:=PostRequest("http://localhost:2222","/rest/annotate",urlValues,"xml")
		fmt.Println("status for entity request", status)

		entitiesMap := make(map[string]int)
		type wrapper struct{
				Resource []Resource `xml:"Resources>Resource"`
		}
		var temp wrapper

		err := xml.Unmarshal( data, &temp)
		if err != nil {
			fmt.Println("Entities xml Unmarshall Error:", err)
		}
		//Need to remove duplicates and get Categories
		entities:=make([]EntityData,0,cap(temp.Resource))

		for _,URI :=range temp.Resource{

			_, prs := entitiesMap[URI.URI]
			if !prs{

				split:=strings.Split(URI.URI,`/`) //Extract term from dbpedia URI
				entity:=split[len(split)-1]
				entitiesMap[entity]+=1 // update the map so that we can ignore if it appears again
				entities=append(entities,GetEntityInfo(entity,"")[0]) // the fist entry

			}
		}

		allData.Entities=entities


}

/*
* Get entities from the Alchemy Api. 
* 
*/
// func GetEntities(allData *ArticleData, linkUrl *string) {
// type Entity struct{
// 	Type string
// 	Relevance string
// 	Count string
// 	Text string
// 	Disambiguated Disambiguatations
// 	Classes []string
// 	Categories []string
// }
// 	apiURL := "http://access.alchemyapi.com/calls/url/URLGetRankedNamedEntities"
// 	fmt.Println("url:= ",apiURL)
// 	urlValues:=url.Values{ "url": {*linkUrl},"apikey": {"7287ca76f8e1b6730787e8ee14ec1b03dacc2041"},"outputMode":{"json"}}
// 	status,data:=PostRequest(apiURL,urlValues)
// 	fmt.Println("status for entity request", status)
// 	//fmt.Println("response Body:", string(data))

// 	type wrapper struct{
// 		Entities []Entity
// 	}

// 	var temp wrapper

// 	err := json.Unmarshal( data, &temp)
// 	if err != nil {
// 		fmt.Println("Entities Json Unmarshall Error:", err)
// 	}


// 	for i:=0;i<len(temp.Entities);i++{

// 		entity:=temp.Entities[i]

// 		results	:= GetEntityInfo(entity.Text,entity.Type)
		
// 		if(len(results)<1){
// 			continue
// 		}

// 		//Dont change to var entity use the index since go is creating copies
// 		temp.Entities[i].Classes=results[0].Classes
// 		temp.Entities[i].Categories=results[0].Categories
// 	}


// 	allData.Entities=temp.Entities
	
// }
	 
/*
* Launch a post request.
 @param url:- the url to which the post request needs to be sents
 @param path, The parth after the base url E.g "/api/weather"
 @param urlValues:- Post requst parameters are passed by creating a url.Values structure e.g:- url.Values{ "param1": "1","param2": "2" }
 @param acceptType:- Specify accept type header if so desired.Only xml has been implmented.Leave blank if so desired
*/

func PostRequest(baseUrl string,path string,urlValues url.Values,acceptType string ) (string, []byte){
	
	Url, err := url.Parse(baseUrl)
	if err != nil {
        panic("invalid Url for get request")
    }
	Url.Path += path
    Url.RawQuery = urlValues.Encode()

	client := &http.Client{}
	fmt.Println(Url.String())
	req, err := http.NewRequest("POST",Url.String(),nil)
	if acceptType=="xml" {
		req.Header.Add("Accept", "text/xml")
	}
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return resp.Status, body

}

/**
Gets classes and Categoried for each entity from the dbpedia server
*/
func GetEntityInfo(name string,entityType string) ([]EntityData){

	values:=url.Values{"QueryClass":{entityType},"QueryString":{name}};
    
	 _,result:=GetRequest("http://localhost:1111","/api/search/KeywordSearch/" ,values)

	temp := EntityDataList{}
	
	xml.Unmarshal(result,&temp)

	for i,entityData:= range temp.EntityDataSlice{
		entityData.Classes=CleanClasses(entityData.Classes)
		temp.EntityDataSlice[i]=entityData
	}
	return temp.EntityDataSlice
	

}
/*
Remove certain classes
*/
var remove=[]string{"owl#Thing"} 

func CleanClasses(classes []string)([]string){
	
	newClasses:=make([]string,0,cap(classes))
	for _,toRemove:= range remove{
		for _,class:= range classes{
			if class!=toRemove{
				newClasses=append(newClasses,class)
			}
		}
	}

	return newClasses
}
