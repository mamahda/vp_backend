package usecase

import (
	"vp_backend/internal/domain"
	"vp_backend/internal/repository"
)

type UserUsecase interface {
	GetUsers() ([]domain.User, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return &userUsecase{repo: r}
}

func (u *userUsecase) GetUsers() ([]domain.User, error) {
	return u.repo.FindAll()
}
