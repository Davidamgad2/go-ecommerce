package routes

import (
	"e-commerce/handlers"

	"github.com/gin-gonic/gin"
)


func InitProductRoutes(router *gin.Engine) {
	productHandler := handlers.ProductHandler{}
	group := router.Group("/products")
	{
		group.GET("", productHandler.FindProductById)
		group.POST("", productHandler.CreateProduct)
		group.PUT("/:id", productHandler.UpdateProduct)
		group.DELETE("/:id", productHandler.DeleteProduct)
	}
}
