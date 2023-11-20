package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ogabrielrodrigues/go-shop/server/internal/model/product"
	"github.com/ogabrielrodrigues/go-shop/server/internal/rest"
)

type productHandler struct{}

func ProductHandler() *productHandler {
	return &productHandler{}
}

func (ph *productHandler) Create(ctx *gin.Context) {
	body := rest.ProductBody{}

	if err := ctx.Bind(&body); err != nil {
		rest.Err(ctx, err)
		return
	}

	new_product, err := product.CreateProduct(&body)
	if err != nil {
		rest.Err(ctx, err)
		return
	}

	rest.Response(ctx, rest.CREATED, new_product.ToResponse())
}

func (ph *productHandler) FindAll(ctx *gin.Context) {
	products, err := product.FindProducts()

	if err != nil {
		rest.Err(ctx, err)
		return
	}

	rest.Response(ctx, rest.OK, products)
}
