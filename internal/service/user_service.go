package service

import (
	"context"

	"vp_backend/internal/repository"
	"vp_backend/internal/domain"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func (s *UserService) GetUser(ctx context.Context, id int) (*domain.User, error) {
	return s.UserRepo.FindByID(ctx, id)
}
