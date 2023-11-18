package product

import (
	"context"

	"github.com/ogabrielrodrigues/go-shop/server/database"
	"github.com/ogabrielrodrigues/go-shop/server/internal/rest"
)

func CreateProduct(body *rest.ProductBody) (*Product, error) {
	db := database.GetPool()

	product := NewProduct(body.Name, body.Brand, body.Description, body.Price, body.Stock)

	_, err := db.Exec(context.Background(), CREATE_QUERY, product.ID, product.Name, product.Brand, product.Description, product.Price, product.Stock)
	if err != nil {
		return nil, database.DatabaseError(err)
	}

	return product, nil
}
