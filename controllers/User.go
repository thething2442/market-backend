package controllers

import (
	"net/http"

	"market-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []models.User

		if err := db.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
			return
		}

		c.JSON(http.StatusOK, users)
	}
}


func GetUserById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
	
		userID := c.Param("userID")
		

		var user models.User

	
		result := db.First(&user, "id = ?", userID)

		
		if result.Error != nil {
	
			if result.Error == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
			}
			return
		}


		c.JSON(http.StatusOK, user)
	}
}