package main

import (
	"log"
	"net/http"
	"time"
	"optiflow/internal/api"
	"optiflow/internal/config"
	"optiflow/internal/db"
	"optiflow/internal/metrics"
	"optiflow/internal/logger"
	"optiflow/pkg/middleware"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize logger
    logger.Init()

	// Load configuration
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	// Initialize metrics
    metrics.Init()

	// Initialize database
	dbConn, err := db.InitDB(cfg)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	// Initialize router
	router := mux.NewRouter()
	router.Handle("/metrics", metrics.MetricsHandler())

	// Middleware
	rl := middleware.NewRateLimiter(100, time.Minute)
	router.Use(middleware.RateLimitMiddleware(rl))
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
