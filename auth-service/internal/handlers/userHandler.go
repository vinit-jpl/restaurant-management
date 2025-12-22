package handlers

import (
	"encoding/json"
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

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid request payload", err)
		return
	}

	user, err := uh.userService.CreateUser(r.Context(), req)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "internal server error", err)
		return
	}

	// send only safe fields
	respData := dto.CreateUserResponse{
		ID:    user.ID,
		Email: user.Email,
	}

	utils.RespondSuccess(
		w,
		http.StatusCreated,
		"user registered",
		respData,
	)
}
