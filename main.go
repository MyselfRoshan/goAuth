package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MyselfRoshan/goAuth/internals/database"
	"github.com/MyselfRoshan/goAuth/internals/routes"
	"github.com/joho/godotenv"
)

func main() {
	// Load the .env File
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Connect to the database
	database.Connect()
	router := routes.Router()
	PORT := fmt.Sprintf(":%v", os.Getenv("PORT"))
	log.Panic(http.ListenAndServe(PORT, router))
}
