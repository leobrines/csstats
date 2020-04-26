package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var db *sql.DB

func main() {
	initDB()
}

func initDB() {
	fmt.Println("Initializing database")

	user := os.Getenv("DB_USER")
	passwd := os.Getenv("DB_PW")
	host := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")

	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", user, passwd, host, dbname))
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected!")
}
