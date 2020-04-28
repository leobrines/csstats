package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

var db *sql.DB

func main() {
	initDB()
	initServer()
}

// initDB create connection to csstats database
func initDB() {
	fmt.Println("Initializing database")

	user := os.Getenv("DB_USER")
	passwd := os.Getenv("DB_PW")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, passwd, host, port, dbname)

	var err error
	db, err = sql.Open("mysql", connection)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected!")
}
