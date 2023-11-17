package routes

import (
	"tmc-api/service"
	"github.com/gin-gonic/gin"
)

func InventoryTransferRoute(router *gin.Engine) {
	router.GET("/tmc/transfer-inventory/get-all", service.GetAllTransfers)
	router.GET("/tmc/transfer-inventory/:id", service.GetTransferInventoryById)
	router.POST("/tmc/transfer-inventory/add-transfer", service.AddTransferInventory)
	router.DELETE("/tmc/transfer-inventory/:id", service.DeleteTransferInventoryById)
	router.PUT("/tmc/transfer-inventory/:id", service.UpdateTransferInventoryById)
	router.POST("/tmc/transfer-inventory/:id/change-inventory-status",service.ChangeInventoryStatus)
}