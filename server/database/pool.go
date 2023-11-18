package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ogabrielrodrigues/go-shop/server/config/env"
)

var pl *pgxpool.Pool

func InitDatabase() (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), env.DATABASE_URL)
	if err != nil {
		return nil, err
	}

	pl = pool

	return pl, nil
}

func GetPool() *pgxpool.Pool {
	return pl
}
