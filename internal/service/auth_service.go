package service

import (
	"errors"
	"time"
	"context"

	"vp_backend/internal/config"
	"vp_backend/internal/domain"
	"vp_backend/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

func (s *AuthService) Register(ctx context.Context, user *domain.User) error {
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password), bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user.Password = string(hashed)
	return s.UserRepo.Create(ctx, user)
}

func (s *AuthService) Login(ctx context.Context, email, password string) (*domain.User ,string, error) {
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
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := token.SignedString([]byte(config.GetJWT()))
	return user, stringToken, err
}

