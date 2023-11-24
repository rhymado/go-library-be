package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware(ctx *gin.Context) {
	// cors diatur oleh beberapa header cors
	// Access-Control-Allow-Origin
	// Access-Control-Allow-Methods
	// Access-Control-Allow-Headers
	// Access-Control-Max-Age
	// Access-Control-Expose-Headers
	// Access-Control-Allow-Credentials
	whitelistOrigin := []string{"http://127.0.0.1:5500", "https://preview-1181115.playcode.io"}
	origin := ctx.GetHeader("Origin")
	for _, worigin := range whitelistOrigin {
		if origin == worigin {
			ctx.Header("Access-Control-Allow-Origin", origin)
			break
		}
	}
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, HEAD, PATCH, PUT, DELETE, OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Authorization")

	// handle preflight
	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}

	ctx.Next()
}
