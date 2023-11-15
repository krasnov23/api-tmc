package service

import (
	"database/sql"
	"net/http"
	"tmc-api/config"
	"tmc-api/models"

	"github.com/gin-gonic/gin"
)

func GetAllInventories(c *gin.Context){
	
	db := config.GetDB()

	rows, err := db.Query("SELECT * FROM inventory")

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	inventories := make([]models.Inventory, 0)

	for rows.Next() {
		var inventory models.Inventory

		err := rows.Scan(&inventory.ID, &inventory.AccountID, &inventory.StateID, &inventory.Ident, &inventory.DatePay, &inventory.DateCreate, &inventory.CategoryID, &inventory.NameID)

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


func GetInventoryById(c *gin.Context) {
	
	var inventory models.Inventory
	id := c.Param("id")

	//log.Println(reflect.TypeOf(id))

	//id,_ = strconv.Atoi(id)

	db := config.GetDB()
	
	row := db.QueryRow("SELECT * FROM inventory WHERE id = $1", id)

	// этот код сканирует возвращенную строку базы данных и связывает значения полей с атрибутами структуры Inventory.
	err := row.Scan(&inventory.ID, &inventory.AccountID, &inventory.StateID, &inventory.Ident, &inventory.DatePay, &inventory.DateCreate, &inventory.CategoryID, &inventory.NameID)
	
	if err != nil {
		if err == sql.ErrNoRows {
			// Если запись не
			c.JSON(404, gin.H{"error": "No row found"})
			return
		} else {
			// при других ошибках, выводим их
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(200, inventory)
}

func AddInventory(c *gin.Context) {
	
	db := config.GetDB()
	
	var inventory models.Inventory
	
	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO inventory(account_id, state_id, ident, date_pay, date_create, category_id, name_id) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := db.QueryRow(query, inventory.AccountID, inventory.StateID, inventory.Ident, inventory.DatePay, inventory.DateCreate, inventory.CategoryID, inventory.NameID).Scan(&inventory.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &inventory)
}

func UpdateInventoryById(c *gin.Context) {
	
	db := config.GetDB()

	getInventorySql := "SELECT * FROM Inventory WHERE id = $1"
	updateInventorySql := "UPDATE Inventory SET account_id=$1, state_id=$2, ident=$3, date_pay=$4, date_create=$5, category_id=$6, name_id=$7 WHERE id = $8"

	inventory := models.Inventory{}
	id := c.Param("id")

	// Находит по id
	row := db.QueryRow(getInventorySql, id)

	// Сопопставляет к структуре Inventory
	err := row.Scan(&inventory.ID, &inventory.AccountID, &inventory.StateID, &inventory.Ident, &inventory.DatePay, &inventory.DateCreate, &inventory.CategoryID, &inventory.NameID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Здесь из тела запроса берется JSON, который преобразовывается обратно в структуру Inventory.
	c.BindJSON(&inventory)

	_, err = db.Exec(
		updateInventorySql,
		inventory.AccountID,
		inventory.StateID,
		inventory.Ident,
		inventory.DatePay,
		inventory.DateCreate,
		inventory.CategoryID,
		inventory.NameID,
		id,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Здесь из тела запроса берется JSON, который преобразовывается обратно в структуру Inventory.
	c.JSON(200, &inventory)
}

func DeleteInventoryById(c *gin.Context) {
	
	db := config.GetDB()

	id := c.Param("id")

	stmt := `DELETE FROM inventory WHERE ID=$1`

	_, err := db.Exec(stmt, id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"success": "inventory with your id has been deleted"})

}