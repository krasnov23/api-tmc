package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect(){
	connStr := "postgres://postgres:postgres@localhost:5436/tmc?sslmode=disable"
	
	// Открывает коннект с базой данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// Тестирует соединение
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	DB = db

	CreateTables()
	fmt.Println("Tables successfully created ")
}


func CheckError(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}

func CreateTables() {
	
	db := GetDB()

	// Create table 'category'
	_, err := db.Exec("CREATE TABLE category IF NOT EXISTS ( id SERIAL PRIMARY KEY, name VARCHAR(50) NOT NULL );")
	CheckError(err)

	// Create table 'inventoryName'
	_, err = db.Exec("CREATE TABLE inventoryName IF NOT EXISTS ( id SERIAL PRIMARY KEY, name VARCHAR(50) NOT NULL, categoryId INTEGER REFERENCES category(id) );")
	CheckError(err)

	// Create table 'inventory'
	_, err = db.Exec("CREATE TABLE inventory IF NOT EXISTS ( id SERIAL PRIMARY KEY, account_id INTEGER, state_id INTEGER, ident VARCHAR(50), date_pay DATE, date_create TIMESTAMP, category_id INTEGER REFERENCES category(id), name_id INTEGER REFERENCES inventoryName(id));")
	CheckError(err)
	
	log.Println("Tables created!")
}


func GetDB() *sql.DB {
	return DB
}