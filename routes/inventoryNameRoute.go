package routes

import (
	"tmc-api/service"

	"github.com/gin-gonic/gin"
)

func InventoryNameRoute(router *gin.Engine) {
	router.GET("/tmc/inventory-name/get-all",service.GetAllInventoriesName)
	router.GET("/tmc/inventory-name/:id", service.GetInventoryNameById)
	router.POST("/tmc/inventory-name/add-name", service.AddInventoryName)
	router.DELETE("/tmc/inventory-name/:id", service.DeleteInventoryNameById)
	router.PUT("/tmc/inventory-name/:id", service.UpdateInventoryNameById)	
}