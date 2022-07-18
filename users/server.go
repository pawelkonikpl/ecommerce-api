package main

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/pawelkonikpl/ecommerce-api/users/database"
	"github.com/pawelkonikpl/ecommerce-api/users/graph"
	"github.com/pawelkonikpl/ecommerce-api/users/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "4001"

func main() {
	port := os.Getenv("USERS_SERVICE_PORT")
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

	DB.AddQueryHook(database.DBLogger{})

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		UsersRepo: database.UsersRepo{DB: DB},
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("🚀 User service is connected to http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
