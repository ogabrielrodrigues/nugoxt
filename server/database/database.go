package database

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ogabrielrodrigues/go-shop/server/config"
)

var Conn *pgxpool.Pool

func NewConnection() (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	Conn, err = pgxpool.Connect(ctx, config.GetDBConfig().ConnString)
	if err != nil {
		return nil, err
	}

	return Conn, nil
}

func DatabaseError(err error) error {
	switch pg_err := err.(*pgconn.PgError); pg_err.Code {
	case "23505":
		column := strings.Split(strings.Split(pg_err.Detail, ")=")[0], "(")[1]
		return fmt.Errorf("query error because `%s` is a unique constraint", column)
	default:
		return err
	}
}
