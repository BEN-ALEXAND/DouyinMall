package controllers

import (
	"dymall/config"
	"dymall/server/models"

	"github.com/gin-gonic/gin"
)

func ListOrders(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(400, gin.H{"error": "user_id is required"})
		return
	}

	var orders []models.Order
	result := config.DB.Where("user_id = ?", userID).Find(&orders)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	
	c.JSON(200, orders)
}

func CreateOrder(c *gin.Context) {
	var newOrder models.Order
	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&newOrder)
	c.JSON(200, newOrder)
}
