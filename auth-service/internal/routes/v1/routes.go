package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/vinit-jpl/restaurant-management/auth-service/internal/handlers"
)

func V1ApiRoutes(userHandler *handlers.UserHandler) chi.Router {
	r := chi.NewRouter()

	r.Get("/test", handlers.Test)
	r.Post("/register", userHandler.CreateUser)

	return r
}
