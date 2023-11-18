package env

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	SERVER_ADDR string
)

func Load() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	SERVER_ADDR = os.Getenv("SERVER_ADDR")

	return nil
}
