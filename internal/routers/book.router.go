package routers

import (
	"lib/internal/handlers"
	"lib/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterBookRouter(router *gin.Engine, db *sqlx.DB) {
	bookRouter := router.Group("/books")
	bookRepository := repositories.InitBookRepository(db)
	bookHandler := handlers.InitBookHandler(bookRepository)
	bookRouter.GET("", bookHandler.GetAllBooks)
}
