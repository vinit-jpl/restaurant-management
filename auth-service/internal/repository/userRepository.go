package repository

import (
	"context"

	"github.com/vinit-jpl/restaurant-management/auth-service/internal/models"
	"gorm.io/gorm"
)

// repository should acccept a Db connection object

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(userDb *gorm.DB) *UserRepository {
	return &UserRepository{db: userDb}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {

	if err := ur.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil

}
