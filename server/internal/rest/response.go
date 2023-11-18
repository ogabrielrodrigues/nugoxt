package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	OK          = http.StatusOK
	BAD_REQUEST = http.StatusBadRequest
)

func Err(ctx *gin.Context, err error) {
	ctx.JSON(BAD_REQUEST, gin.H{
		"error": err,
	})
}
