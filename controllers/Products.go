// controllers/product.go

package controllers

import (
	"net/http"

	"market-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetProducts retrieves all products from the database.
func GetProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []models.Product
		if err := db.Find(&products).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
			return
		}
		c.JSON(http.StatusOK, products)
	}
}

// GetProductById retrieves a single product by its ID from the URL.
func GetProductById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		productID := c.Param("productID")
		var product models.Product
		if err := db.First(&product, "id = ?", productID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product"})
			}
			return
		}
		c.JSON(http.StatusOK, product)
	}
}