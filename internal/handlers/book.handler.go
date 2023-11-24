package handlers

import (
	"lib/internal/models"
	"lib/internal/repositories"
	"log"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	*repositories.BookRepository
}

func InitBookHandler(rep *repositories.BookRepository) *BookHandler {
	return &BookHandler{rep}
}

func (b *BookHandler) GetAllBooks(ctx *gin.Context) {
	// auth := ctx.GetHeader("Authorization")
	// fmt.Println(auth)
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

func (b *BookHandler) CreateBook(ctx *gin.Context) {
	body := &models.BookModel{}
	if err := ctx.ShouldBind(body); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if _, err := govalidator.ValidateStruct(body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err := b.CreateNewBook(body); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Buku berhasil dibuat",
	})
}
