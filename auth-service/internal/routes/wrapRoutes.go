package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/vinit-jpl/restaurant-management/auth-service/internal/handlers"
	v1 "github.com/vinit-jpl/restaurant-management/auth-service/internal/routes/v1"
)

func WrapRoutes(userHandler *handlers.UserHandler) chi.Router {

	r := chi.NewRouter()

	// Pass handler to V1 routes
	r.Mount("/v1", v1.V1ApiRoutes(userHandler))

	return r
}
