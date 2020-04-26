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
	dbname := os.Getenv("DB_NAME")
	connection := fmt.Sprintf("%s:%s@%s/%s", user, passwd, host, dbname)

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
