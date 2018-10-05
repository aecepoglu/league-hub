package main

import (
	"log"
	"net/http"
	"io/ioutil"
	//"path/filepath"
	//"os"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func foo() int {
	return 2
}

func main() {
	b, err := ioutil.ReadFile("schema.graphql")

	if err != nil {
		log.Fatal(err)
		return
	}

	str := string(b)
	//log.Printf("Serving files in %s\n", assets)
	schema := graphql.MustParseSchema(str, &resolvers{})
	http.Handle("/graphql", &relay.Handler {Schema: schema})
	http.Handle("/", http.FileServer(http.Dir("./assets")))

	log.Println("Running at :8080");
	err = http.ListenAndServe(":8080", nil)

	log.Fatal(err)
}
