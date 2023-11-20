package product

import (
	"time"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/go-shop/server/internal/rest"
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

func (p *Product) ToResponse() *rest.ProductResponse {
	return &rest.ProductResponse{
		ID:          p.ID.String(),
		Name:        p.Name,
		Brand:       p.Brand,
		Description: p.Description,
		Price:       p.Price,
		Images:      p.Images,
		Stock:       p.Stock,
		CreatedAt:   p.CreatedAt,
	}
}
