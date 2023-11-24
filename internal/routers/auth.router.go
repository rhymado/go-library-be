package routers

import (
	"lib/internal/handlers"
	"lib/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterAuthRouter(router *gin.Engine, db *sqlx.DB) {
	authRouter := router.Group("/auth")
	authRepository := repositories.InitAuthRepository(db)
	authHandler := handlers.InitAuthHandler(authRepository)

	authRouter.POST("", authHandler.LoginHandler)
	authRouter.POST("/new", authHandler.RegisterHandler)
}
