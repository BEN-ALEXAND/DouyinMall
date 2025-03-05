package main

import (
	"log"

	"dymall/config"
	"dymall/server/routes"
)

func main() {
	config.InitDB()
	// config.InitWechatPay()

	// set up router
	r := routes.SetupRouter()

	// start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}