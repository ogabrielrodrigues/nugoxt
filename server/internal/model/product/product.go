package product

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID
	Name        string
	Brand       string
	Description string
	Price       float64
	Images      []string
	Stock       uint64
	CreatedAt   time.Time
}

func NewProduct(name, brand, description string, price float64, stock uint64) *Product {
	return &Product{
		ID:          uuid.New(),
		Name:        name,
		Brand:       brand,
		Description: description,
		Price:       price,
		Images:      []string{},
		Stock:       stock,
		CreatedAt:   time.Now(),
	}
}
