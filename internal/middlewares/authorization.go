package middlewares

import (
	"lib/pkg"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTGate(allowedRole ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := ctx.GetHeader("Authorization")
		if bearerToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Please login first",
			})
			return
		}

		if !strings.Contains(bearerToken, "Bearer") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Please login again",
			})
			return
		}

		token := strings.Replace(bearerToken, "Bearer ", "", -1)
		payload, err := pkg.VerifyToken(token)
		if err != nil {
			if strings.Contains(err.Error(), "expired") {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "Please login again",
				})
				return
			}
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		var allowed = false
		for _, role := range allowedRole {
			if payload.Role == role {
				allowed = true
				break
			}
		}

		if !allowed {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Maaf akses tidak diperbolehkan",
			})
			return
		}
		ctx.Set("Payload", payload)
		ctx.Next()
	}
}
