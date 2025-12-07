package main

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vinit-jpl/restaurant-management/auth-service/internal/config"
	"github.com/vinit-jpl/restaurant-management/auth-service/internal/database/postgres"
	"github.com/vinit-jpl/restaurant-management/auth-service/internal/models"
	"github.com/vinit-jpl/restaurant-management/auth-service/internal/routes"
	"github.com/vinit-jpl/restaurant-management/auth-service/internal/server"
)

func main() {

	cfg := config.LoadConfig()

	dbDriver := postgres.NewPostgresDb(cfg)
	db, _ := dbDriver.Connect()

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Database Migration Failed: %v", err)
	}

	log.Println("Database Migrated Successfully")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)

	//register routes
	r.Mount("/authservice/api", routes.WrapRoutes())
	// Initialize and start the server

	srv := server.NewServer(cfg.Port, r)
	log.Println("Auth service up and running")
	srv.Start()

}
