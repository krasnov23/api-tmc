package main

import (
	"tmc-api/routes"

	"github.com/gin-gonic/gin"
	"tmc-api/config"
)

func main() {
	server := gin.Default()

	config.Connect()
	
	routes.InventoryRoute(server)
	routes.InventoryNameRoute(server)
	routes.InventoryCategoryRoute(server)

	server.Run(":8080")

}