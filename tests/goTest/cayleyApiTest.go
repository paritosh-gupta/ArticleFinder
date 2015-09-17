package main

import _ "github.com/google/cayley/graph/bolt"
import (
	"fmt"
	"github.com/google/cayley"
	// "github.com/google/cayley/graph"
	//"github.com/google/cayley/writer"
)

const dbPath="/Users/guptap/Documents/SearchTool/Database/backend.db"

func main() {
	 //graph.InitQuadStore("bolt", dbPath, nil)
	store, err := cayley.NewGraph("bolt", dbPath, nil)
	if err != nil {
		fmt.Println("error in creating database", err)
	}
	err=store.AddQuad(cayley.Quad("food", "is", "good", ""))
	// if err != nil {
	// 	fmt.Println("write error", err)
	// }
	err=store.AddQuad(cayley.Quad("nothing", "is", "good", ""))
}
