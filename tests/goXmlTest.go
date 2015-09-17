package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml"
)

type EntityData struct{
	Name string `xml:"Label"`
	Classes []string `xml:"Classes>Class>Label"`
	Categories []string `xml:"Categories>Category>Label"`
}

func main() {
	_,result:=GetRequest("http://localhost:1111/api/search/KeywordSearch?QueryString=xfinity")

	temp := EntityDataList{}
	
	xml.Unmarshal(result,&temp)
	
	fmt.Println(cap(temp.EntityDataSlice))
	// fmt.Println(temp.Results[0].Classes[0])
	fmt.Println(temp)
}

type EntityDataList struct{
	 EntityDataSlice []EntityData `xml:"Result"`
}

func GetRequest(url string) (string, []byte) {

	fmt.Println("Get request to URL:>", url)
	req, err := http.NewRequest("GET", url, nil)
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
