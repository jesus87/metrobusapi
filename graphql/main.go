package main

import (
	"log"
	"net/http"
	"os"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)
//main main of the api, and opens the port for using api on local
func main() {

	port := os.Getenv("GRAPHQL_PORT")

	router := fasthttprouter.New()

	graphqlHandler := NewGraphqlHandler()

	router.Handle(http.MethodGet, "/", graphqlHandler.Handle)

	log.Fatal(fasthttp.ListenAndServe(port, router.Handler))
}
