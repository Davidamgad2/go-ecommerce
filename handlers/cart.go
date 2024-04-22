package handlers

import (
	"e-commerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CartHandler struct {
}

func (CartHandler) CreateCart(c *gin.Context) {
	// Create a new cart
	cart := models.Cart.NewCart(models.Cart{})
	db := c.MustGet("db").(*gorm.DB)
	isSucceeded := db.Create(&cart)
	if isSucceeded.Error != nil {
	c.JSON(http.StatusOK, cart)
	return
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
}

func (h CartHandler) FindCartById(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var cart models.Cart
	db.First(&cart, c.Param("id"))
	if cart.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}
	c.JSON(http.StatusOK, cart)
}

func (h CartHandler) UpdateCart(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var cart models.Cart
	if err := db.First(&cart, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if err := db.Save(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart"})
		return
	}
	c.JSON(http.StatusOK, cart)
}

func (h CartHandler) DeleteCart(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var cart models.Cart
	if err := db.First(&cart, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}
	if err := db.Delete(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete cart"})
		return
	}
	c.Status(http.StatusOK)
}
