package routes

import (
	"tmc-api/service"

	"github.com/gin-gonic/gin"
)

func InventoryCategoryRoute(router *gin.Engine) {
	router.GET("/tmc/inventory-category/get-all",service.GetAllCategories)
	router.GET("/tmc/inventory-category/:id", service.GetCategoryById)
	router.POST("/tmc/inventory-category/add", service.AddCategory)
	router.DELETE("/tmc/inventory-category/:id", service.DeleteCategory)
	router.PUT("/tmc/inventory-category/:id", service.UpdateCategory)	
}

