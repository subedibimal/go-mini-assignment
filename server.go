package main

import (
	"log"
	"net/http"
	"os"

	"github.com/subedibimal/go-mini-assignment/config"
	"github.com/subedibimal/go-mini-assignment/directives"
	"github.com/subedibimal/go-mini-assignment/graph"
	"github.com/subedibimal/go-mini-assignment/graph/generated"
	"github.com/subedibimal/go-mini-assignment/middlewares"
	"github.com/subedibimal/go-mini-assignment/migration"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
)

const defaultPort = "8080"

func main() {
	migration.MigrateTable()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := config.GetDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	router := mux.NewRouter()
	router.Use(middlewares.AuthMiddleware)

	c := generated.Config{Resolvers: &graph.Resolver{}}
	c.Directives.Auth = directives.Auth

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
