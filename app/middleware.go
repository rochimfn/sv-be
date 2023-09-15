package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitMiddleware(r *gin.Engine) *gin.Engine {
	r.Use(cors.Default())

	return r
}
