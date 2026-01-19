package service

import (
	"context"

	"vp_backend/internal/domain"
	"vp_backend/internal/repository"
)

type PropertyService struct {
	PropertyRepo *repository.PropertyRepository
}

func (s *PropertyService) Create(ctx context.Context, p *domain.Property) error {
	return s.PropertyRepo.Create(ctx, p)
}

func (s *PropertyService) GetByID(ctx context.Context, id int) (*domain.Property, error) {
	return s.PropertyRepo.FindByID(ctx, id)
}

func (s *PropertyService) GetAll(ctx context.Context) ([]domain.Property, error) {
	return s.PropertyRepo.FindAll(ctx)
}

func (s *PropertyService) GetFilteredProperty(ctx context.Context, f *domain.PropertyFilters) ([]domain.Property, int, error) {
	if f.Page <= 0 {
		f.Page = 1
	}
	if f.Limit <= 0 {
		f.Limit = 10
	}
	if f.Limit > 100 {
		f.Limit = 100
	}

	f.Offset = (f.Page - 1) * f.Limit

	return s.PropertyRepo.FindFiltered(ctx, f)
}

func (s *PropertyService) Update(ctx context.Context, p *domain.Property) error {
	return s.PropertyRepo.Update(ctx, p)
}

func (s *PropertyService) Delete(ctx context.Context, id int) error {
	return s.PropertyRepo.Delete(ctx, id)
}
