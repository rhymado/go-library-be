package routers

import (
	"lib/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterUploadRouter(router *gin.Engine, db *sqlx.DB) {
	uploadRouter := router.Group("/upload")
	uploadHandler := handlers.InitUploadHandler()

	uploadRouter.POST("", uploadHandler.UploadToCloud)
	uploadRouter.POST("/local", uploadHandler.SavedToDirectory)
}
