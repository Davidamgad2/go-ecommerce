package routes

import (
	"e-commerce/handlers"

	"github.com/gin-gonic/gin"
)

func InitCartRoutes(router *gin.Engine) {
	cartHandler := handlers.CartHandler{} // Create an instance of CartHandler
	group := router.Group("/carts")
	{
		group.GET("", cartHandler.FindCartById)
		group.POST("", cartHandler.CreateCart)
		group.PUT("/:id", cartHandler.UpdateCart)
		group.DELETE("/:id", cartHandler.DeleteCart)
	}
}
