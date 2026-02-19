package service

import (
	"context"

	"vp_backend/internal/domain"
	"vp_backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// UserService menangani business logic
// yang berkaitan dengan data dan profil user.
type UserService struct {
	UserRepo *repository.UserRepository
}

// Get mengambil data user berdasarkan ID.
func (s *UserService) Get(
	ctx context.Context,
	id int,
) (*domain.User, error) {

	return s.UserRepo.FindByID(ctx, id)
}

// UpdateUser memperbarui data profil user.
//
// Alur:
// 1. Ambil data user berdasarkan ID
// 2. Perbarui field yang diizinkan
// 3. Simpan perubahan ke database
func (s *UserService) UpdateUser(
	ctx context.Context,
	id int,
	username string,
	email string,
	phone string,
	password string,
) error {

	user, err := s.UserRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user.Username = username
	user.Email = email
	user.Phone = phone
	user.Password = string(hashedPassword)

	return s.UserRepo.Update(ctx, user)
}

