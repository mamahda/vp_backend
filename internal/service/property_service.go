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

func (s *PropertyService) Update(ctx context.Context, p *domain.Property) error {
	return s.PropertyRepo.Update(ctx, p)
}

func (s *PropertyService) Delete(ctx context.Context, id int) error {
	return s.PropertyRepo.Delete(ctx, id)
}

