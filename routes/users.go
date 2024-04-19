package routes

import (
	"github.com/gin-gonic/gin"
)

func UserHandler(c *gin.Context) {
	// Handle User logic here
}

func InitUserRoutes(router *gin.Engine) {
	group := router.Group("/users")
	{
		group.GET("", UserHandler)
		group.POST("", UserHandler)
		group.PUT("/:id", UserHandler)
		group.DELETE("/:id", UserHandler)
	}
}
