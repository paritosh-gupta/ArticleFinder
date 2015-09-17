package extractArticleData
/*
* This package is responsible 
	1) Extracting article text,title,etc.
		This is done by calling the python serverwhich runs goose.
	2) Calling the DBpedia Spotlight Server to extract entities.
	3) Calling the DBpedia lookup Server to extract information for each entity
	
	See README for more information on each server
	
	Terminology:-
	An Article is the link that a user gives to be input into the archive

	
*/


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

const gooseServerAddress = "http://localhost:8081" // Address of the server that extracts article data(text,title,image etc)
const dbpediaSpotlightAddress = "http://localhost:2222" //Dbpedia server which is used for entity extraction
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
*Information Extracted for each entity
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
This function extracts all information from an Article.
@param linkUrl string:=The link of the article
return struct of type ArticleData with all the information
*/
func GetDataFromArticle(linkUrl string) ArticleData {
	
	var allData ArticleData
	ExtractArticleTextInfo(&allData, &linkUrl)
	GetEntities(&allData)
	allData.Parent="Article" //The parent node of every link is Article
	allData.Link=linkUrl
	allData.Date=time.Now().Local().Format("2006-01-02 15:04:05 +0800")
	return allData
	
}

/**
* Extracts article information from the python Goose server() .
  @param allData *ArticleData : pointer to an articleData to where all information needs to be stored
*/
func ExtractArticleTextInfo(allData *ArticleData, linkUrl *string) {
	values:=url.Values{"url": {*linkUrl}}
	_, data := GetRequest(gooseServerAddress,"",values)
	err := json.Unmarshal(data, &allData)
	if err != nil {
		fmt.Println("Error in Extract Article Info json:", err)
	}
}

/*
	Launch a Get Request. No Headers have been set
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

/*
*Struct to Parse Data Returned by the DBpedia Spotlight server
*/
type Resource struct{
			URI string `xml:"URI,attr"`
}

/*
* Get Entities from the DBPedia Spotlight server.
* @param allData ArticleData:- An ArticleData Instance which must have the article contents
in allData.Content to get some result
*/
func GetEntities(allData *ArticleData){
		
		urlValues := url.Values{}
		urlValues.Set("text", allData.Content)
		urlValues.Add("powered_by", "no")
		urlValues.Add("policy", "whitelist")
		urlValues.Add("confidence", "0.5")
		status,data:=PostRequest("http://10.2.10.52:2222","/rest/annotate",urlValues,"xml")
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
		//Need to remove duplicates and get Categories.Hence a map
		entities:=make([]EntityData,0,cap(temp.Resource))

		for _,URI :=range temp.Resource{

			_, valuePresent := entitiesMap[URI.URI]
			if !valuePresent{

				split:=strings.Split(URI.URI,`/`) //Extract term from dbpedia URI

				entity:=split[len(split)-1]
				fmt.Println("entity is "+entity)
				entitiesMap[URI.URI]+=1 // update the map so that we can ignore if it appears again
				entityData:=GetEntityInfo(entity,"")
				if len(entityData)>0{
					entities=append(entities,GetEntityInfo(entity,"")[0]) // the first entry
				}
			}
		}

		allData.Entities=entities


}
	 
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
	// fmt.Println(Url.String())
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
Gets classes and Categories for each entity from the dbpedia server
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
