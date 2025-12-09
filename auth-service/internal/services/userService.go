package services

import (
	"context"

	"github.com/vinit-jpl/restaurant-management/auth-service/internal/dto"
	"github.com/vinit-jpl/restaurant-management/auth-service/internal/models"
	"github.com/vinit-jpl/restaurant-management/auth-service/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{repo: userRepo}
}

func (us *UserService) CreateUser(ctx context.Context, req dto.CreateUserRequest) (*models.User, error) {

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Phone:    req.Phone,
	}

	createdUser, err := us.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
