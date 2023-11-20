package product

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/go-shop/server/database"
	"github.com/ogabrielrodrigues/go-shop/server/internal/rest"
)

func CreateProduct(body *rest.ProductBody) (*Product, error) {
	db := database.Conn
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	product := NewProduct(body.Name, body.Brand, body.Description, body.Price, body.Stock)

	_, err := db.Exec(ctx, CREATE_QUERY, product.ID, product.Name, product.Brand, product.Description, product.Price, product.Stock)
	if err != nil {
		return nil, database.DatabaseError(err)
	}

	return product, nil
}

func FindProducts() (*[]Product, error) {
	db := database.Conn
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := db.Query(ctx, FIND_ALL_QUERY)
	if err != nil {
		return nil, database.DatabaseError(err)
	}

	products := []Product{}

	for rows.Next() {
		row, _ := rows.Values()

		price, _ := strconv.ParseFloat(fmt.Sprintf("%g", row[4].(float32)), 64)

		product := Product{
			ID:          uuid.MustParse(row[0].(string)),
			Name:        row[1].(string),
			Brand:       row[2].(string),
			Description: row[3].(string),
			Price:       price,
			Images:      []string{},
			Stock:       uint64(row[5].(int16)),
			CreatedAt:   row[6].(time.Time),
		}

		products = append(products, product)
	}

	return &products, nil
}
