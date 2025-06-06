package main

import (
	"fp-sbd/db"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.ConnectMongoDB()
	db.ConnectMySQL()
}
