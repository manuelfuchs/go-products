package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	load_environment()

	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DATABASE"))

	a.Run(8010)
}

func load_environment() {
	// Load the .env file
	error := godotenv.Load("../.env")
	if error != nil {
		log.Fatal("Error loading .env file")
	}
}
