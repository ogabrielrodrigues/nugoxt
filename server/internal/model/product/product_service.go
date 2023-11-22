package product

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/ogabrielrodrigues/go-shop/server/database"
	"github.com/ogabrielrodrigues/go-shop/server/internal/rest"
)

const (
	CREATE_QUERY        = `INSERT INTO tb_products (id, name, brand, description, price, stock, created_at) VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP)`
	CREATE_IMAGES_QUERY = `INSERT INTO tb_images (product_id, url) VALUES ($1, $2)`
	FIND_ALL_QUERY      = `SELECT tb_products.*, array_agg(tb_images.url) AS images FROM tb_products JOIN tb_images ON tb_products.id = tb_images.product_id GROUP BY tb_products.id`
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

	batch := &pgx.Batch{}
	for _, img := range body.Images {
		batch.Queue(CREATE_IMAGES_QUERY, product.ID, img)
	}

	_, err = db.SendBatch(ctx, batch).Exec()
	if err != nil {
		return nil, err
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
		var images []string

		for _, img := range row[7].(pgtype.TextArray).Elements {
			images = append(images, img.String)
		}

		product := Product{
			ID:          uuid.MustParse(row[0].(string)),
			Name:        row[1].(string),
			Brand:       row[2].(string),
			Description: row[3].(string),
			Price:       price,
			Images:      images,
			Stock:       uint64(row[5].(int16)),
			CreatedAt:   row[6].(time.Time),
		}

		products = append(products, product)
	}

	return &products, nil
}
