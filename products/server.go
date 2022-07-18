package main

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/pawelkonikpl/ecommerce-api/products/database"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/pawelkonikpl/ecommerce-api/products/graph"
	"github.com/pawelkonikpl/ecommerce-api/products/graph/generated"
)

const defaultPort = "4002"

func main() {
	port := os.Getenv("PRODUCTS_SERVICE_PORT")
	if port == "" {
		port = defaultPort
	}

	DB := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "example",
		Database: "ecommerce_api_users",
	})

	defer func(DB *pg.DB) {
		err := DB.Close()
		if err != nil {
			fmt.Printf("Cannot close DB connection: %s", err)
		}
	}(DB)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		ProductsRepo: database.ProductsRepo{
			DB: DB,
		},
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("🚀 Product service is connected to http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
