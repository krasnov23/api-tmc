package service

import (
	"database/sql"
	"net/http"
	"strconv"
	"tmc-api/config"
	"tmc-api/models"

	"github.com/gin-gonic/gin"
)

func GetAllTransfers(c *gin.Context){

	db := config.GetDB()

	rows, err := db.Query("SELECT * FROM inventory_transfer")

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	transfers := make([]models.InventoryTransfer, 0)

	for rows.Next() {
		var transfer models.InventoryTransfer

		err := rows.Scan(&transfer.ID,&transfer.SenderID,&transfer.ReceiverID,&transfer.Ident,&transfer.TransferDate,&transfer.Status)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		transfers = append(transfers, transfer)
	}

	if err = rows.Err(); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, &transfers)
}


func AddTransferInventory(c *gin.Context) {
	db := config.GetDB()

	var transfer models.InventoryTransfer

	if err := c.ShouldBindJSON(&transfer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO inventory_transfer(sender_id,reciever_id,ident,transfer_date,status) VALUES($1, $2, $3, $4, $5) RETURNING id`
	err := db.QueryRow(query, transfer.SenderID,transfer.ReceiverID,transfer.Ident,transfer.TransferDate,transfer.Status).Scan(&transfer.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &transfer)
}

func GetTransferInventoryById(c *gin.Context){
	var transfer models.InventoryTransfer
	
	id := c.Param("id")

	db := config.GetDB()
	
	row := db.QueryRow("SELECT * FROM inventory_transfer WHERE id = $1", id)

	// этот код сканирует возвращенную строку базы данных и связывает значения полей с атрибутами структуры Inventory.
	err := row.Scan(&transfer.ID,&transfer.SenderID,&transfer.ReceiverID,&transfer.Ident,&transfer.TransferDate,&transfer.Status)
	
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

	c.JSON(200, transfer)

}

func UpdateTransferInventoryById(c *gin.Context){
	
	db := config.GetDB()

	transfer := models.InventoryTransfer{}
	err := c.BindJSON(&transfer)
	if err != nil {
		c.JSON(400, gin.H{"error": "Bad request"})
		return
	}

	
	sqlStatement := `UPDATE Inventory_transfer SET sender_id = $1, reciever_id = $2, ident = $3, 
	transfer_date = $4, status = $5 WHERE id = $6`
	_, err = db.Exec(sqlStatement,transfer.SenderID, transfer.ReceiverID, transfer.Ident, transfer.TransferDate, transfer.Status, c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update inventory name"})
		return
	}

	transfer.ID,_ = strconv.Atoi(c.Param("id"))

	c.JSON(200, &transfer)
}

func DeleteTransferInventoryById(c *gin.Context){
	
	db := config.GetDB()

	stmt := `DELETE FROM inventory_transfer WHERE ID=$1`

	_, err := db.Exec(stmt, c.Param("id"))

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"success": "inventory Transfer with your id has been deleted"})
}

func ChangeInventoryStatus(c *gin.Context){
	
	var transfer models.InventoryTransfer
	var inventory models.Inventory

	id := c.Param("id")

	db := config.GetDB()
	
	row := db.QueryRow("SELECT * FROM inventory_transfer WHERE id = $1", id)

	// этот код сканирует возвращенную строку базы данных и связывает значения полей с атрибутами структуры Inventory.
	err := row.Scan(&transfer.ID,&transfer.SenderID,&transfer.ReceiverID,&transfer.Ident,&transfer.TransferDate,&transfer.Status)
	
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

	var statusUpdate struct {
        Status string `json:"status"`
    }
    
	if err := c.ShouldBindJSON(&statusUpdate); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	
	if statusUpdate.Status == "в обработке" || statusUpdate.Status == "отклонен"{
		
		query := `UPDATE inventory_transfer SET status = $1 WHERE id = $2`
    	_, err = db.Exec(query, statusUpdate.Status, id)
    	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    	}
		
		transfer.Status = statusUpdate.Status

		c.JSON(200, &transfer)

	}else if statusUpdate.Status == "подтвержден"{
		
		query := `UPDATE inventory_transfer SET status = $1 WHERE id = $2`
    	
		_, err = db.Exec(query, statusUpdate.Status, id)
    	if err != nil {
        
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    	}

		_ = db.QueryRow("UPDATE Inventory SET account_id=$1,date_create=$2 WHERE ident = $3", transfer.ReceiverID,transfer.TransferDate,transfer.Ident)

		getUpdateInventory := db.QueryRow("SELECT * FROM Inventory WHERE ident = $1", transfer.Ident)

		err := getUpdateInventory.Scan(&inventory.ID, &inventory.AccountID, &inventory.StateID, &inventory.Ident, &inventory.DatePay, &inventory.DateCreate, &inventory.CategoryID, &inventory.NameID)
		
		//c.BindJSON(&inventory)

		if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
		}

		c.JSON(200, &inventory)
	}else{
		c.JSON(500, gin.H{"error": "несуществующий статус"})
	}
	
	
}