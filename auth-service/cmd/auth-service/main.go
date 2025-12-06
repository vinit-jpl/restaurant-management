package main

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vinit-jpl/restaurant-management/auth-service/internal/config"
	"github.com/vinit-jpl/restaurant-management/auth-service/internal/server"
)

func main() {

	cfg := config.LoadConfig()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Initialize and start the server

	srv := server.NewServer(cfg.Port, r)
	log.Println("Auth service up and running")
	srv.Start()

}
