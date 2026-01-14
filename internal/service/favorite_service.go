package service

import (
	"context"
	"vp_backend/internal/domain"
	"vp_backend/internal/repository"
)

type FavoriteService struct {
	FavoriteRepo *repository.FavoriteRepository
}

func (s *FavoriteService) AddFavorite(ctx context.Context, userID int, propertyID int) error {
	return s.FavoriteRepo.Add(ctx, userID, propertyID)
}

func (s *FavoriteService) RemoveFavorite(ctx context.Context, userID int, propertyID int) error {
	return s.FavoriteRepo.Remove(ctx, userID, propertyID)
}

func (s *FavoriteService) GetAll(ctx context.Context, userID int) ([]domain.Property, error) {
	return s.FavoriteRepo.FindAll(ctx, userID)
}
