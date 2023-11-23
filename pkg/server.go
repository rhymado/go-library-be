package pkg

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func InitServer(router *gin.Engine) *http.Server {
	var addr string = "localhost:8080"
	server := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 10,
		Handler:      router,
	}
	return server
}
