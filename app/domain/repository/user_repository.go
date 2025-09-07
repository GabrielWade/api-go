package repository

import "api/app/domain/models"

type UserRepository interface {
	GetAll() ([]models.User, error)
	Create(user *models.User) error
	FindByUsername(username string) (*models.User, error)
}
