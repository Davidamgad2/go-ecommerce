package routes

import (
	"e-commerce/handlers"

	"github.com/gin-gonic/gin"
)

func UserHandler(c *gin.Context) {
	// Handle User logic here
}

func InitUserRoutes(router *gin.Engine) {
	userHandler := handlers.UserHandler{}
	group := router.Group("/users")
	{
		group.GET("", userHandler.FindUserById)
		group.POST("", userHandler.CreateUser)
		group.PUT("/:id", userHandler.UpdateUser)
		group.DELETE("/:id", userHandler.DeleteUser)
	}
}
