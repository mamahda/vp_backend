package service

import (
	"context"

	"vp_backend/internal/repository"
	"vp_backend/internal/domain"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func (s *UserService) Get(ctx context.Context, id int) (*domain.User, error) {
	return s.UserRepo.FindByID(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, id int, username, email, phone string) error {
	user, err := s.UserRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	user.Username = username
	user.Email = email
	user.Phone = phone

	return s.UserRepo.Update(ctx, user)
}
