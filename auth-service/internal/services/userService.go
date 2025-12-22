package services

import (
	"context"

	"github.com/vinit-jpl/restaurant-management/auth-service/internal/dto"
	"github.com/vinit-jpl/restaurant-management/auth-service/internal/models"
	"github.com/vinit-jpl/restaurant-management/auth-service/internal/repository"
	"github.com/vinit-jpl/restaurant-management/auth-service/internal/utils"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{repo: userRepo}
}

func (us *UserService) CreateUser(ctx context.Context, req dto.CreateUserRequest) (*models.User, error) {

	hashedPassword, err := utils.Hashpassword(req.Password)

	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Phone:    req.Phone,
	}

	createdUser, err := us.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
