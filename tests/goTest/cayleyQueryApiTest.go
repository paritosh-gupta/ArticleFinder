package main

import _ "github.com/google/cayley/graph/bolt"
import (
	"fmt"
	"github.com/google/cayley"
	 "github.com/google/cayley/graph"
	// _ "github.com/google/cayley/writer"
	// "github.com/google/cayley/quad"
	// "log"
)

const dbPath = "/Users/guptap/Documents/SearchTool/Database/backend.db"

func main() {

	store, err := cayley.NewGraph("bolt", dbPath, nil)
	if err != nil {
		fmt.Println("error in creating database", err)
	}

	path := cayley.StartPath(store, "Article").
		In().
		Tag("link").
		Save("has_image", "image").
		Save("has_title","title").
		Save("has_description","description")

	it := path.BuildIterator()
	it, _ = it.Optimize()

	for graph.Next(it) {
		tags := make(map[string]graph.Value)
		it.TagResults(tags)
		fmt.Println( store.NameOf(tags["image"]))
		fmt.Println( store.NameOf(tags["title"]))
		fmt.Println( store.NameOf(tags["description"]))
		fmt.Println( store.NameOf(tags["link"]))
	}

}
