package rest

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	OK          = http.StatusOK
	CREATED     = http.StatusCreated
	BAD_REQUEST = http.StatusBadRequest
)

type R map[string]interface{}

type ProductResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Brand       string    `json:"brand"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Images      []string  `json:"images"`
	Stock       uint64    `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
}

func Err(ctx *gin.Context, err error) {
	ctx.JSON(BAD_REQUEST, gin.H{
		"error": err.Error(),
	})
}

func Response(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, data)
}
