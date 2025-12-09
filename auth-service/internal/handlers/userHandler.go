package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/vinit-jpl/restaurant-management/auth-service/internal/dto"
	"github.com/vinit-jpl/restaurant-management/auth-service/internal/services"
	"github.com/vinit-jpl/restaurant-management/auth-service/internal/utils"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(us *services.UserService) *UserHandler {
	return &UserHandler{userService: us}
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest

	// Decode JSON body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, errors.New("invalid request body"))
		return
	}

	// Service call
	user, err := uh.userService.CreateUser(r.Context(), req)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err)
		return
	}

	// Success response
	utils.RespondSuccess(w, http.StatusCreated, "user created successfully", dto.CreateUserResponse{
		ID:    user.ID,
		Email: user.Email,
	})
}
