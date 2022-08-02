package database

import (
	"context"
	"github.com/go-pg/pg/v10"
	"log"
)

type DBLogger struct{}

func (d DBLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d DBLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	query, _ := q.FormattedQuery()
	log.Println("USER SERVICE", string(query))
	return nil
}
