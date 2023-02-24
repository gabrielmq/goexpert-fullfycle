package database

import (
	"github.com/gabrielmq/apis/internal/entity"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{db}
}

func (u *User) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *User) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	query := u.DB.Where("email = ?", email)
	if err := query.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
