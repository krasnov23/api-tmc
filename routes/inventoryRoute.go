package routes

import (
	"github.com/gin-gonic/gin"
	"tmc-api/service"
)

func InventoryRoute(router *gin.Engine) {
	router.GET("/tmc/inventory/get-all",service.GetAllInventories)
	router.GET("/tmc/inventory/:id", service.GetInventoryById)
	router.POST("/tmc/inventory/add-inventory", service.AddInventory)
	router.DELETE("/tmc/inventory/:id", service.DeleteInventoryById)
	router.PUT("/tmc/inventory/:id", service.UpdateInventoryById)
}

