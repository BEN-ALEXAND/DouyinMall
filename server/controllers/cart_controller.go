package controllers

import (
	"dymall/config"
	"dymall/server/models"

	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCart(c *gin.Context) {
	var newCart models.Cart
	if err := c.ShouldBindJSON(&newCart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&newCart)
	c.JSON(200, newCart)
}

func DisplayCart(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(400, gin.H{"error": "user_id is required"})
		return
	}

	var cart []models.Cart
	result := config.DB.Where("user_id = ?", userID).Find(&cart)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	
	c.JSON(200, cart)
}

// Clean the cart
func ClearCart(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(400, gin.H{"error": "user_id is required"})
		return
	}

	result := config.DB.Where("user_id = ?", userID).Delete(&models.Cart{})
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
	}
}