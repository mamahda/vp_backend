package service

import (
	"context"
	// "mime/multipart"
	// "net/http"

	"vp_backend/internal/domain"
	"vp_backend/internal/repository"
	"vp_backend/internal/storage"
)

// PropertyService menangani business logic
// yang berkaitan dengan data properti.
type PropertyService struct {
	PropertyRepo *repository.PropertyRepository
	Storage      storage.Storage
}

// Create menyimpan data properti baru ke database.
func (s *PropertyService) Create(
	ctx context.Context,
	p *domain.Property,
) error {

	return s.PropertyRepo.Create(ctx, p)
}


func (s *PropertyService) GetPropertyImages(
	ctx context.Context,
	id int,
) ([]string, error) {

	return s.PropertyRepo.FindPropertyImagesUrl(ctx, id)
}

// GetByID mengambil detail properti
// berdasarkan ID properti.
func (s *PropertyService) GetByID(
	ctx context.Context,
	id int,
) (*domain.Property, error) {

	return s.PropertyRepo.FindByID(ctx, id)
}

// GetAll mengambil seluruh data properti
// tanpa filter atau pagination.
func (s *PropertyService) GetAll(
	ctx context.Context,
) ([]domain.Property, error) {

	return s.PropertyRepo.FindAll(ctx)
}

func (s *PropertyService) GetCountData(
	ctx context.Context,
	f *domain.PropertyFilters,
) (int, error) {
	return s.PropertyRepo.CountData(ctx, f)
}

// GetFilteredProperty mengambil daftar properti
// berdasarkan filter, sorting, dan pagination.
//
// Validasi yang dilakukan:
// - Page minimal 1
// - Limit default 10
// - Limit maksimal 100
func (s *PropertyService) GetFilteredProperty(
	ctx context.Context,
	f *domain.PropertyFilters,
) ([]domain.Property, error) {

	if f.Page <= 0 {
		f.Page = 1
	}
	if f.Limit <= 0 {
		f.Limit = 10
	}
	if f.Limit > 100 {
		f.Limit = 100
	}

	// Hitung offset untuk pagination
	f.Offset = (f.Page - 1) * f.Limit

	return s.PropertyRepo.FindFiltered(ctx, f)
}

// Update memperbarui data properti
// berdasarkan ID.
func (s *PropertyService) Update(
	ctx context.Context,
	p *domain.Property,
) error {

	return s.PropertyRepo.Update(ctx, p)
}

// Delete menghapus data properti
// berdasarkan ID.
func (s *PropertyService) Delete(
	ctx context.Context,
	id int,
) error {

	return s.PropertyRepo.Delete(ctx, id)
}
