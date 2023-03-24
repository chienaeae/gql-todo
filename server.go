package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/chienaeae/gql-todo/graph"
	"github.com/chienaeae/gql-todo/internal/auth"
	"github.com/chienaeae/gql-todo/internal/directives"
	database "github.com/chienaeae/gql-todo/internal/pkg/db/migrations/mysql"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(auth.Middleware())

	database.InitDB()
	defer database.CloseDB()
	database.Migrate()

	c := graph.Config{Resolvers: &graph.Resolver{}}
	c.Directives.Binding = directives.Binding
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
