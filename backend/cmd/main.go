package main

import (
	"backend/internal/database"
	"log"
)

func main() {
	//init db
	database.InitDB()
	defer database.CloseDB()

	log.Println("Server is running...")
}