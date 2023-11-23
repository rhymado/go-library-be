package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterRouter(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	RegisterBookRouter(router, db)
	return router
}
