package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/MalukiMuthusi/edufi/server/db"
	"github.com/MalukiMuthusi/edufi/server/graph"
	"github.com/MalukiMuthusi/edufi/server/graph/generated"
)

const defaultPort = "8110"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	p := db.Postgres{}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{S: &p}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func init() {
	db.StartDB()
}
