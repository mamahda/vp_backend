package service

import (
	"context"
	"errors"
	"time"

	"vp_backend/internal/config"
	"vp_backend/internal/domain"
	"vp_backend/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// AuthService berisi business logic
// yang berkaitan dengan autentikasi user,
// seperti register dan login.
type AuthService struct {
	UserRepo *repository.UserRepository
}

// Register melakukan proses registrasi user baru.
//
// Alur:
// 1. Password di-hash menggunakan bcrypt
// 2. Data user disimpan ke database melalui repository
//
// Error yang mungkin:
// - error dari bcrypt
// - domain.ErrEmailAlreadyExists (dari repository)
func (s *AuthService) Register(
	ctx context.Context,
	user *domain.User,
) error {

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	return s.UserRepo.Create(ctx, user)
}

// Login memverifikasi kredensial user dan
// menghasilkan JWT token jika berhasil.
//
// Alur:
// 1. Cari user berdasarkan email
// 2. Bandingkan password plaintext dengan hash bcrypt
// 3. Generate JWT token dengan expired 24 jam
//
// Return:
// - data user
// - JWT token (string)
// - error jika gagal
func (s *AuthService) Login(
	ctx context.Context,
	email string,
	password string,
) (*domain.User, string, error) {

	user, err := s.UserRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	); err != nil {
		return nil, "", errors.New("invalid password")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	stringToken, err := token.SignedString(
		[]byte(config.GetJWT()),
	)

	return user, stringToken, err
}
