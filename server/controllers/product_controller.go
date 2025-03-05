package controllers

import (
	"dymall/config"
	"dymall/server/models"

	"github.com/gin-gonic/gin"
)

func ListProducts(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)
	c.JSON(200, products)
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&product)
	c.JSON(200, product)
}