package handlers

import (
	"fmt"
	"lib/internal/repositories"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	*repositories.BookRepository
}

func InitBookHandler(rep *repositories.BookRepository) *BookHandler {
	return &BookHandler{rep}
}

func (b *BookHandler) GetAllBooks(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")
	fmt.Println(auth)
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	result, err := b.ReadAllBooks(page, limit)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": "Success",
	})
}
