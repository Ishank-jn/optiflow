package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"optiflow/api"
	"optiflow/config"
	"optiflow/db"
	"optiflow/middleware"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	// Initialize database
	dbConn, err := db.InitDB(cfg)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	// Initialize router
	router := mux.NewRouter()

	// Middleware
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.AuthMiddleware)

	// API Handlers
	api.SetupRoutes(router, dbConn)

	// Start server
	log.Printf("Starting server on %s", cfg.Server.Address)
	if err := http.ListenAndServe(cfg.Server.Address, router); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
