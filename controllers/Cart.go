// controllers/cart.go

package controllers

import (
	"net/http"

	"market-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetCarts retrieves all carts from the database.
func GetCarts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var carts []models.Cart
		if err := db.Find(&carts).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve carts"})
			return
		}
		c.JSON(http.StatusOK, carts)
	}
}

// GetCartById retrieves a single cart by its ID from the URL.
func GetCartById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		cartID := c.Param("cartID")
		var cart models.Cart
		if err := db.First(&cart, "id = ?", cartID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cart"})
			}
			return
		}
		c.JSON(http.StatusOK, cart)
	}
}