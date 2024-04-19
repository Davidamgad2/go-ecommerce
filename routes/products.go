package routes

import (
	"github.com/gin-gonic/gin"
)

func ProductHandler(c *gin.Context) {
	// Handle Product logic here
}

func InitProductRoutes(router *gin.Engine) {
	group := router.Group("/products")
	{
		group.GET("", ProductHandler)
		group.POST("", ProductHandler)
		group.PUT("/:id", ProductHandler)
		group.DELETE("/:id", ProductHandler)
	}
}
