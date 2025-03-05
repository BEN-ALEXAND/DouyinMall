package routes

import (
	"dymall/server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// users router
	usersGroup := router.Group("/users")
	{
		usersGroup.POST("/register", controllers.Register)
		usersGroup.POST("/login", controllers.Login)
	}

	// products router
	productsGroup := router.Group("/products")
	{
		productsGroup.GET("/ProductList", controllers.ListProducts)
		productsGroup.POST("/CreateProduct", controllers.CreateProduct)
	}

	// orders router
	ordersGroup := router.Group("/orders")
	{
		ordersGroup.GET("/OrderList", controllers.ListOrders)
		ordersGroup.POST("/CreateOrder", controllers.CreateOrder)
	}

	// cart router
	cartGroup := router.Group("/cart")
	{
		cartGroup.GET("/", controllers.DisplayCart)
	}

	// payment router
	paymentGroup := router.Group("/payment")
	{
		paymentGroup.POST("/", controllers.ProcessPayment)
	}

	return router
}