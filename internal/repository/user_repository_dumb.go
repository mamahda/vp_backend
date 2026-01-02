package repository

import "vp_backend/internal/domain"

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindAll() ([]domain.User, error) {
	return []domain.User{
		{ID: 1, Name: "Alice", Email: "alice@mail.com"},
		{ID: 2, Name: "Bob", Email: "bob@mail.com"},
	}, nil
}
