package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ogabrielrodrigues/go-shop/server/config/env"
	"github.com/ogabrielrodrigues/go-shop/server/database"
	"github.com/ogabrielrodrigues/go-shop/server/internal/router"
)

func main() {
	server := gin.Default()
	gp := server.Group("/api/v1")

	if err := env.Load(); err != nil {
		panic("error on load environment variables")
	}

	pool, err := database.InitDatabase()
	if err != nil {
		panic("error on initialize database pool")
	}

	defer pool.Close()

	router.InitRouter(gp)

	server.Run(env.SERVER_ADDR)
}
