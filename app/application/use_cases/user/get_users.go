package user

import (
    "api/app/domain/models"
    "api/app/domain/repository"
)

type GetUsersUseCase struct {
    repo repository.UserRepository
}

func NewGetUsersUseCase(repo repository.UserRepository) *GetUsersUseCase {
    return &GetUsersUseCase{repo}
}

func (uc *GetUsersUseCase) Execute() ([]models.User, error) {
    return uc.repo.GetAll()
}