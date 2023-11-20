package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect(){
	
	godotenv.Load()
	
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL_MODE"),
	)
	
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

	_,err := db.Exec("ALTER TABLE  IF EXISTS InventoryName RENAME TO inventory_name;")
	CheckError(err)

	_,err = db.Exec("ALTER TABLE  IF EXISTS InventoryTransfer RENAME TO inventory_transfer;")
	CheckError(err)

	// Create table 'category'
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS category ( id SERIAL PRIMARY KEY, name VARCHAR(50) NOT NULL );")
	CheckError(err)

	// Create table 'inventoryName'
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS inventory_name ( id SERIAL PRIMARY KEY, name VARCHAR(50) NOT NULL, categoryId INTEGER REFERENCES category(id) );")
	CheckError(err)

	// Create table 'inventory'
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS inventory ( id SERIAL PRIMARY KEY, account_id INTEGER, state_id INTEGER, ident VARCHAR(50), date_pay DATE, date_create TIMESTAMP, category_id INTEGER REFERENCES category(id), name_id INTEGER REFERENCES inventoryName(id));")
	CheckError(err)
	
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS inventory_transfer  ( id SERIAL PRIMARY KEY, sender_id INTEGER, reciever_id INTEGER, ident VARCHAR(50), transfer_date TIMESTAMP, status VARCHAR(50));")
	CheckError(err)

	log.Println("Tables created!")
}


func GetDB() *sql.DB {
	return DB
}