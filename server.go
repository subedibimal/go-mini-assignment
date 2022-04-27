package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/subedibimal/go-mini-assignment/graph"
	"github.com/subedibimal/go-mini-assignment/graph/generated"
	"github.com/subedibimal/go-mini-assignment/migration"
	"github.com/subedibimal/go-mini-assignment/config"
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

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
