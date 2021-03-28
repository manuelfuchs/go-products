package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func load_environment() {
	// Load the .env file
	error := godotenv.Load(".env")
	if error != nil {
		log.Fatal("Error loading .env file")
	}
}

func (a *App) InitializeFromEnvironment() {
	load_environment()

	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))
}

func (a *App) Initialize(user, password, database string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, database)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

func (a *App) Run(port uint16) {

}
