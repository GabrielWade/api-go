package user

import (
	"api/app/domain/models"
	"api/app/domain/repository"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserUseCase struct {
	repo repository.UserRepository
}

func NewCreateUserUseCase(repo repository.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{repo}
}

func (uc *CreateUserUseCase) Execute(user *models.User) error {
	hash, err := generatePasswordHash(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash
	return uc.repo.Create(user)
}

func generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
