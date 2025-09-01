package main

import (
	"log"
	"net/http"

	"go-backend-demo/internal/config"
	"go-backend-demo/internal/database"
	"go-backend-demo/internal/handlers"
	"go-backend-demo/pkg/middleware"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := database.Connect(cfg.Database)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Initialize handlers
	userHandler := handlers.NewUserHandler(db)

	//create router
	r := mux.NewRouter()

	//add middleware
	r.Use(middleware.LoggingMiddleware)

	//api routes
	api := r.PathPrefix("/api/v1").Subrouter()

	//user routes
	api.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	api.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	api.HandleFunc("/users/{id:[0-9]+}", userHandler.GetUser).Methods("GET")
	api.HandleFunc("/users/{id:[0-9]+}", userHandler.UpdateUser).Methods("PUT")
	api.HandleFunc("/users/{id:[0-9]+}", userHandler.DeleteUser).Methods("DELETE")

	// Health check
	api.HandleFunc("/health", userHandler.HealthCheck).Methods("GET")

	// Start server
	log.Printf("Server starting on port %s\n", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(cfg.Server.Port, r))
}
