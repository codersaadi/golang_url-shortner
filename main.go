package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
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
	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.Logger)
	// router.Use(cors.Handler(cors.Options{
	// 	// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
	// 	AllowedOrigins:   []string{"https://*", "http://*"},
	// 	// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
	// 	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	// 	ExposedHeaders:   []string{"Link"},
	// 	AllowCredentials: false,
	// 	MaxAge:           300, // Maximum value not ignored by any of major browsers
	//   }))

	// Register routes
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
		log.Fatal("DB_URL not set in env")
	}
	conn, err := sql.Open("postgres", DB_URL)
	if err != nil {
		log.Fatal("error opening postgres", err)
	}
	queries := database.New(conn)
	apiCfg := app.ApiConfig{
		DB: queries,
	}

	return apiCfg
}
