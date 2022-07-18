package graph

import "github.com/pawelkonikpl/ecommerce-api/users/database"

//go:generate go run github.com/99designs/gqlgen generate

// migrate create -ext sql -dir migrations -seq create_users_table
// migrate -database ${POSTGRESQL_URL} -path db/migrations up

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	database.UsersRepo
}
