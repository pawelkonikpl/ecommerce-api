package database

import "github.com/go-pg/pg/v10"

type UsersRepo struct {
	DB *pg.DB
}
