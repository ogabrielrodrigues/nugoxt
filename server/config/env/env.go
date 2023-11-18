package env

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	SERVER_ADDR  string
	DATABASE_URL string
)

func Load() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	SERVER_ADDR = os.Getenv("SERVER_ADDR")
	DATABASE_URL = os.Getenv("DATABASE_URL")

	return nil
}
