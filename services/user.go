package services

import (
	"forum-api/config"
	"forum-api/models"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService() *UserService {
	return &UserService{db: config.DB}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.db.Create(user).Error
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := s.db.Where("email = ?", email).First(&user).Error
	return &user, err
}
