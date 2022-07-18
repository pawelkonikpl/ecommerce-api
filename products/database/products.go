package database

import "github.com/go-pg/pg/v10"

type ProductsRepo struct {
	DB *pg.DB
}
