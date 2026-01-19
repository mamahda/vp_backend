package service

import (
	"context"

	"vp_backend/internal/domain"
	"vp_backend/internal/repository"
)

// FavoriteService berisi business logic
// yang berkaitan dengan fitur favorit properti user.
type FavoriteService struct {
	FavoriteRepo *repository.FavoriteRepository
}

// AddFavorite menambahkan properti ke daftar favorit user.
//
// Parameter:
// - userID     : ID user yang sedang login
// - propertyID : ID properti yang ingin difavoritkan
func (s *FavoriteService) AddFavorite(
	ctx context.Context,
	userID int,
	propertyID int,
) error {

	return s.FavoriteRepo.Add(ctx, userID, propertyID)
}

// RemoveFavorite menghapus properti dari daftar favorit user.
//
// Parameter:
// - userID     : ID user yang sedang login
// - propertyID : ID properti yang ingin dihapus dari favorit
func (s *FavoriteService) RemoveFavorite(
	ctx context.Context,
	userID int,
	propertyID int,
) error {

	return s.FavoriteRepo.Remove(ctx, userID, propertyID)
}

// GetAll mengambil seluruh daftar properti
// yang telah difavoritkan oleh user.
func (s *FavoriteService) GetAll(
	ctx context.Context,
	userID int,
) ([]domain.Property, error) {

	return s.FavoriteRepo.FindAll(ctx, userID)
}

