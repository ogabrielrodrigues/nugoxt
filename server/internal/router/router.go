package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ogabrielrodrigues/go-shop/server/internal/handler"
)

func InitRouter(gp *gin.RouterGroup) {
	ph := handler.ProductHandler()

	gp.POST("/products", ph.Create)
	gp.GET("/products", ph.FindAll)
}
