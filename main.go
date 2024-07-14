package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // This is the missing import for the PostgreSQL driver
	"github.com/saadi925/url-shortner-golang/internal/app"
	"github.com/saadi925/url-shortner-golang/internal/database"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}
func main() {

	apiCfg := dbInit()

	// Set up router
	router := mux.NewRouter()
	app.RegisterRoutes(router, apiCfg)

	// Start server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
func dbInit() app.ApiConfig {
	DB_URL := os.Getenv("DB_URL")
	if DB_URL == "" {
		log.Fatal("DB_URL not is in env")
	}
	conn, err := sql.Open("postgres", DB_URL)
	if err != nil {
		log.Fatal("error opening postgres ", err)
	}
	queries := database.New(conn)
	apiCfg := app.ApiConfig{
		DB: queries,
	}

	return apiCfg

}
