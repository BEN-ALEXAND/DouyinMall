package controllers

import (
	"dymall/config"
	"dymall/server/models"

	"github.com/gin-gonic/gin"

	"strconv"
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

func SearchProducts(c *gin.Context) {
	var products []models.Product
	productName := c.Query("product_name")
	minPriceStr := c.Query("min_price")
	maxPriceStr := c.Query("max_price")

	var minPrice, maxPrice float64
	var err error

	if minPriceStr != "" {
		minPrice, err = strconv.ParseFloat(minPriceStr, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid price_min value"})
			return
		}
	}

	if maxPriceStr != "" {
		maxPrice, err = strconv.ParseFloat(maxPriceStr, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid price_max value"})
			return
		}
	}

	db := config.DB.Model(&models.Product{})

	if productName != "" {
		db = db.Where("name LIKE ?", "%" + productName + "%")
	}

	if minPriceStr != "" {
		db = db.Where("price >= ?", minPrice)
	}

	if maxPriceStr != "" {
		db = db.Where("price <= ?", maxPrice)
	}

	result := db.Find(&products)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, products)
}