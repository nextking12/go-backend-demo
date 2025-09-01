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

}