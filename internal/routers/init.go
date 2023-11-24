package routers

import (
	"lib/internal/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterRouter(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.CORSMiddleware)

	RegisterBookRouter(router, db)
	return router
}
