package routes

import (
	"github.com/go-chi/chi/v5"
	v1 "github.com/vinit-jpl/restaurant-management/auth-service/internal/routes/v1"
)

func WrapRoutes() chi.Router {

	r := chi.NewRouter()

	r.Mount("/v1", v1.V1ApiRoutes())

	return r
}
