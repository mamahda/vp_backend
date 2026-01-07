package service

import (
	"vp_backend/internal/repository"
	"vp_backend/internal/domain"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func (s *UserService) GetUser(id int) (*domain.User, error) {
	return s.UserRepo.FindByID(id)
}
