package user

import (
	"api/app/domain/repository"
	"api/app/shared/utils/jwt"

	"golang.org/x/crypto/bcrypt"
)

type LoginUserUseCase struct {
	repo repository.UserRepository
}

func NewLoginUserUseCase(repo repository.UserRepository) *LoginUserUseCase {
	return &LoginUserUseCase{repo}
}

func (uc *LoginUserUseCase) Execute(username, password string) (string, string, error) {
	user, err := uc.repo.FindByUsername(username)
	if err != nil {
		return "", "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", err
	}

	accessTokenJWT, err := jwt.GenerateAccessToken(user.ID, user.Username)
	if err != nil {
		return "", "", err
	}

	refreshTokenJWT, err := jwt.GenerateRefreshToken(user.ID, user.Username)
	if err != nil {
		return "", "", err
	}
	return accessTokenJWT, refreshTokenJWT, nil
}
