package service

import (
	"database/sql"
	"net/http"
	"strconv"
	"tmc-api/config"
	"tmc-api/models"
	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context){
	
	db := config.GetDB()

	rows, err := db.Query("SELECT * FROM category")

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	categories := make([]models.Category, 0)

	for rows.Next() {
		var category models.Category

		err := rows.Scan(&category.ID, &category.Name)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, &categories)
}


func GetCategoryById(c *gin.Context) {

	var category models.Category
	id := c.Param("id")

	db := config.GetDB()

	row := db.QueryRow("SELECT * FROM category WHERE id = $1", id)

	// этот код сканирует возвращенную строку базы данных и связывает значения полей с атрибутами структуры Category.
	err := row.Scan(&category.ID, &category.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			// При ошибке No rows - означает что категория не найдена
			c.JSON(404, gin.H{"error": "No row found"})
			return
		} else {
			// при других ошибках, выводим их
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(200, category)
}

func AddCategory(c *gin.Context){
	db := config.GetDB()
	
	var category models.Category
	
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO category(name) VALUES($1) RETURNING id`
	
	err := db.QueryRow(query, category.Name).Scan(&category.ID)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &category)
}

func UpdateCategory(c *gin.Context){
	db := config.GetDB()

	category := models.Category{}
	err := c.BindJSON(&category)
	if err != nil {
		c.JSON(400, gin.H{"error": "Bad request"})
		return
	}

	sqlStatement := `UPDATE category SET name=$1 WHERE id=$2`
	_, err = db.Exec(sqlStatement, category.Name, c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update inventory name"})
		return
	}

	category.ID,_ = strconv.Atoi(c.Param("id"))

	c.JSON(200, &category)
}

func DeleteCategory(c *gin.Context){
	db := config.GetDB()

	id := c.Param("id")

	stmt := `DELETE FROM category WHERE ID=$1`

	_, err := db.Exec(stmt, id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"success": "inventory name with your id has been deleted"})
}




