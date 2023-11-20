package service

import (
	"database/sql"
	"net/http"
	"strconv"
	"tmc-api/config"
	"tmc-api/models"

	"github.com/gin-gonic/gin"
)

func GetAllInventoriesName(c *gin.Context){
	
	db := config.GetDB()

	rows, err := db.Query("SELECT * FROM inventory_name")

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	inventories := make([]models.InventoryName, 0)

	for rows.Next() {
		var inventory models.InventoryName

		err := rows.Scan(&inventory.ID, &inventory.Name,&inventory.CategoryId)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		inventories = append(inventories, inventory)
	}

	if err = rows.Err(); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, &inventories)
}


func GetInventoryNameById(c *gin.Context) {

	var inventoryName models.InventoryName
	id := c.Param("id")

	db := config.GetDB()

	row := db.QueryRow("SELECT * FROM inventory_name WHERE id = $1", id)

	// этот код сканирует возвращенную строку базы данных и связывает значения полей с атрибутами структуры Inventory.
	err := row.Scan(&inventoryName.ID, &inventoryName.Name,&inventoryName.CategoryId)
	
	if err != nil {
		if err == sql.ErrNoRows {
			// Если запись не найдена
			c.JSON(404, gin.H{"error": "No row found"})
			return
		} else {
			// при других ошибках, выводим их
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(200, inventoryName)
}

func AddInventoryName(c *gin.Context) {
	
	db := config.GetDB()
	
	var inventoryName models.InventoryName
	
	if err := c.ShouldBindJSON(&inventoryName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO inventory_name(name,categoryId) VALUES($1, $2) RETURNING id`
	
	err := db.QueryRow(query, inventoryName.Name,inventoryName.CategoryId).Scan(&inventoryName.ID)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &inventoryName)
}

func UpdateInventoryNameById(c *gin.Context){
	
	db := config.GetDB()

	inventoryName := models.InventoryName{}
	err := c.BindJSON(&inventoryName)
	if err != nil {
		c.JSON(400, gin.H{"error": "Bad request"})
		return
	}

	
	sqlStatement := `UPDATE inventory_name SET name=$1, categoryId=$2 WHERE id=$3`
	_, err = db.Exec(sqlStatement, inventoryName.Name, inventoryName.CategoryId, c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update inventory name"})
		return
	}

	inventoryName.ID,_ = strconv.Atoi(c.Param("id"))

	c.JSON(200, &inventoryName)
}

func DeleteInventoryNameById(c *gin.Context) {
	
	db := config.GetDB()

	id := c.Param("id")

	stmt := `DELETE FROM inventory_name WHERE ID=$1`

	_, err := db.Exec(stmt, id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"success": "inventory name with your id has been deleted"})

}








