package handlers

import (
	"e-commerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductHandler struct {
}

func (h ProductHandler) CreateProduct(c *gin.Context) {
	// Create a new product
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}
	c.JSON(http.StatusCreated, product)
}

func (h ProductHandler) FindProductById(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var product models.Product
	if err := db.First(&product, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h ProductHandler) UpdateProduct(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var product models.Product
	if err := db.First(&product, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if err := db.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h ProductHandler) DeleteProduct(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var product models.Product
	if err := db.First(&product, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	if err := db.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}
	c.Status(http.StatusOK)
}
