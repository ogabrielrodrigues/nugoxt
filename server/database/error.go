package database

import (
	"fmt"
	"strings"

	"github.com/jackc/pgconn"
)

func DatabaseError(err error) error {
	switch pg_err := err.(*pgconn.PgError); pg_err.Code {
	case "23505":
		column := strings.Split(strings.Split(pg_err.Detail, ")=")[0], "(")[1]
		return fmt.Errorf("query error because `%s` is a unique constraint", column)
	default:
		return err
	}
}
