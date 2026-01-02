package repository

import "vp_backend/internal/domain"

type UserRepository interface {
	FindAll() ([]domain.User, error)
}
