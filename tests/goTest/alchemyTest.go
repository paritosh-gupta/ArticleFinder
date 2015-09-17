package main

import (
		"net/http"
  		"fmt"
		"io/ioutil"
 		"net/url"
)

type LinkData struct {
	Link   string
	Date   string
	Parent string
}

func post() http.Response {

	apiURL := "http://access.alchemyapi.com/calls/url/URLGetRankedNamedEntities"
	fmt.Println("url:= ",apiURL)

	 resp, err := http.PostForm(apiURL,
	 	url.Values{ "url": {"http://adage.com/article/cmo-strategy/pizza-hut-official-pizza-espn-fantasy-football/300094/"},"apikey": {"7287ca76f8e1b6730787e8ee14ec1b03dacc2041"}})

	if err != nil {
		fmt.Println("invalid error")
		return
	}

	defer resp.Body.Close()
	
	// body, err := ioutil.ReadAll(resp.Body)
	
	// if nil != err {
 //   	 fmt.Println("errorination happened reading the body", err)
 //    return
 //  }
  	return resp
	//fmt.Println(string(body[:]))
	
}

func main(){

	post()
}
