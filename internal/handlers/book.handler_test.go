package handlers

import (
	"encoding/json"
	"lib/internal/models"
	"lib/internal/repositories"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var brm = repositories.BookRepositoryMock{}
var handler = InitBookHandler(&brm)

type Response struct {
	Data    []models.BookResponseModel `json:"data"`
	Message string                     `json:"message"`
}

func TestGetAllBooks(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	res := httptest.NewRecorder()

	r.GET("/books", handler.GetAllBooks)
	ex := []models.BookResponseModel{}
	brm.On("ReadAllBooks", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(ex, nil)

	req := httptest.NewRequest("GET", "/books", nil)
	r.ServeHTTP(res, req)

	exRes := &Response{
		Data:    ex,
		Message: "Success",
	}
	b, err := json.Marshal(exRes)
	if err != nil {
		t.Fatalf("Marshal Error: %e", err)
	}

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, string(b), res.Body.String())
}

func TestCreateBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	res := httptest.NewRecorder()

	r.POST("/books", handler.CreateBook)
	brm.On("CreateNewBook", mock.Anything).Return(nil)

	body := &models.BookModel{
		BookName:    "Books",
		Price:       10000,
		AuthorId:    1,
		PublisherId: 1,
		PromoId:     1,
	}
	b, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Marshal Error: %e", err)
	}

	req := httptest.NewRequest("POST", "/books", strings.NewReader(string(b)))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(res, req)

	exRes := gin.H{
		"status":  "OK",
		"message": "Buku berhasil dibuat",
	}
	bres, err := json.Marshal(exRes)
	if err != nil {
		t.Fatalf("Marshal Error: %e", err)
	}

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, string(bres), res.Body.String())
}
