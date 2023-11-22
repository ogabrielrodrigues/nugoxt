package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ogabrielrodrigues/go-shop/server/config"
	"github.com/ogabrielrodrigues/go-shop/server/database"
	"github.com/ogabrielrodrigues/go-shop/server/internal/router"
)

func main() {
	err := config.Load()
	if err != nil {
		panic(err)
	}

	server := gin.Default()
	gp := server.Group("/api/v1")

	conn, err := database.NewConnection()
	if err != nil {
		panic("error on initialize database pool")
	}

	defer conn.Close()

	router.InitRouter(gp)

	server.Run(config.GetAPIConfig().Port)
}
