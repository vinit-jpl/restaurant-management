package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:255;not null" json:"name"`
	Email    string `gorm:"size:255;not null;uniqueIndex" json:"email"`
	Password string `gorm:"size:255;not null" json:"-"`
	Role     string `gorm:"size:100;default:user" json:"role"`
	Phone    string `gorm:"not null" json:"phone"`
}
