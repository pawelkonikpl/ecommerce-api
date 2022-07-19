package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-pg/pg/v10"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pawelkonikpl/ecommerce-api/users/database"
	"github.com/pawelkonikpl/ecommerce-api/users/graph"
	"github.com/pawelkonikpl/ecommerce-api/users/graph/generated"
	"log"
	"net/http"
	"os"
)

const defaultPort = "4001"

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

	DB := pg.Connect(&pg.Options{
		User:     dbUser,
		Password: dbPassword,
		Database: dbDatabase,
	})

	m, err := migrate.New(
		"file://database/migrations",
		"postgres://postgres:example@localhost:5432/ecommerce_api?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}

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
