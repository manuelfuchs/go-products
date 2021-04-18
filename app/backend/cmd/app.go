package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

const (
	HOST = "db"
	PORT = 5432
)

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
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, user, password, database)
	fmt.Println("----------------------- debug -----------------------")
	fmt.Println(connectionString)
	fmt.Println("----------------------- debug -----------------------")

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(port uint16) {
	var address string = fmt.Sprintf(":%d", port)
	fmt.Printf("Starting with address %s\n", address)

	log.Fatal(http.ListenAndServe(address, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/products", a.getProducts).Methods("GET")
	a.Router.HandleFunc("/product", a.createProduct).Methods("POST")
	a.Router.HandleFunc("/product/{id:[0-9]+}", a.getProduct).Methods("GET")
	a.Router.HandleFunc("/product/{id:[0-9]+}", a.updateProduct).Methods("PUT")
	a.Router.HandleFunc("/product/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")
	a.Router.HandleFunc("/products/below/{maxPrice:[0-9]+}", a.getProductsBelowPrice).Methods("GET")
	a.Router.HandleFunc(
		"/products/between/{minPrice:[0-9]+}/and/{maxPrice:[0-9]+}",
		a.getProductsInPriceRange).Methods("GET")
	a.Router.HandleFunc("/products/containing/{searchtext}", a.getProductsBySearchtext).Methods("GET")
}

func (a *App) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	p := product{ID: id}
	if err := p.getProduct(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Product not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func (a *App) createProduct(w http.ResponseWriter, r *http.Request) {
	var p product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := p.createProduct(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
}

func (a *App) getProducts(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	products, err := getProducts(a.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func (a *App) updateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var p product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	p.ID = id

	if err := p.updateProduct(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func (a *App) deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	p := product{ID: id}
	if err := p.deleteProduct(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *App) getProductsBelowPrice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	maxPrice, err := strconv.Atoi(vars["maxPrice"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid maximum price")
		return
	}

	products, err := getProductsBelowPrice(a.DB, maxPrice)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func (a *App) getProductsInPriceRange(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	minPrice, err := strconv.Atoi(vars["minPrice"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid min price")
		return
	}

	maxPrice, err := strconv.Atoi(vars["maxPrice"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid max price")
		return
	}

	products, err := getProductsInPriceRange(a.DB, minPrice, maxPrice)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func (a *App) getProductsBySearchtext(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	searchtext := vars["searchtext"]
	if len(searchtext) == 0 {
		respondWithError(w, http.StatusBadRequest, "Invalid search text")
		return
	}

	products, err := getProductsBySearchtext(a.DB, searchtext)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
