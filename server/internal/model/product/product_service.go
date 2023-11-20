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

		// 	img_rows, err := db.Query(ctx, FIND_ALL_IMAGES_OF_PRODUCT_QUERY, product.ID)
		// 	if err != nil {
		// 		return nil, err
		// 	}

		// 	var images []string
		// 	for img_rows.Next() {
		// 		img_row, _ := img_rows.Values()

		// 		images = append(images, img_row[0].(string))
		// 	}

		// 	product.Images = images

		products = append(products, product)
	}

	return &products, nil
}
