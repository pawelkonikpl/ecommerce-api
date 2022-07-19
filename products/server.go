package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-pg/pg/v10"
	"github.com/golang-migrate/migrate/v4"
	"github.com/pawelkonikpl/ecommerce-api/products/database"
	"github.com/pawelkonikpl/ecommerce-api/products/graph"
	"github.com/pawelkonikpl/ecommerce-api/products/graph/generated"
	"log"
	"net/http"
	"os"
)

const defaultPort = "4002"

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	port := getenv("PORT", defaultPort)
	dbDatabase := getenv("PG_DB", "ecommerce_api")
	dbPassword := getenv("PG_PASSWORD", "example")
	//dbPort := getenv("PG_PORT", "5432")
	dbUser := getenv("PG_USER", "postgres")

	m, err := migrate.New(
		"file://database/migrations",
		"postgres://postgres:example@localhost:5432/ecommerce_api?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
	DB := pg.Connect(&pg.Options{
		User:     dbUser,
		Password: dbPassword,
		Database: dbDatabase,
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
