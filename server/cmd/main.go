package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ogabrielrodrigues/go-shop/server/config/env"
	"github.com/ogabrielrodrigues/go-shop/server/internal/router"
)

func main() {
	server := gin.Default()
	gp := server.Group("/api/v1")

	if err := env.Load(); err != nil {
		panic("error on load environment variables")
	}

	router.InitRouter(gp)

	server.Run(env.SERVER_ADDR)
}
