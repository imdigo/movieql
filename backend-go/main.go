package main

import (
	"io/ioutil"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/imdigo/movieql/backend-go/graphql"
)

const schemaPath string = "./graphql/schema.graphql"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type query struct{}

func (*query) Hello() string { return "Hello, world!" }

func main() {
	dat, err := ioutil.ReadFile(schemaPath)
	check(err)
	// fmt.Print(string(dat))
	GetMovies()
	s := string(dat)
	schema := graphql.MustParseSchema(s, &query{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
