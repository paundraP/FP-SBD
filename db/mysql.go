package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var MySQL *sql.DB

func ConnectMySQL() {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		log.Fatal("MYSQL_DSN not set")
	}

	var err error
	MySQL, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to open MySQL connection: %v", err)
	}

	if err := MySQL.Ping(); err != nil {
		log.Fatalf("MySQL ping failed: %v", err)
	}

	log.Println("Connected to MySQL")
}
