package main

import (
	"log"
	"net/http"
	"os"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func main() {

	port := os.Getenv("GRAPHQL_PORT")

	router := fasthttprouter.New()

	graphqlHandler := NewGraphqlHandler()

	router.Handle(http.MethodGet, "/", graphqlHandler.Handle)

	log.Fatal(fasthttp.ListenAndServe(port, router.Handler))
}
