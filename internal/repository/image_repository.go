package repository

import (
	"context"
	"database/sql"
	"errors"
	"vp_backend/internal/domain"
	// "vp_backend/internal/domain"
)

// PropertyRepository bertanggung jawab untuk
// mengelola seluruh operasi database yang
// berkaitan dengan entitas Property.
type ImageRepository struct {
	DB *sql.DB
}

func (r *ImageRepository) UpdateCoverImage(ctx context.Context, id int, url string) error {
	query := `UPDATE properties SET cover_image_url = ? WHERE id = ?`
	_, err := r.DB.ExecContext(ctx, query, url, id)
	return err
}

func (r *ImageRepository) SaveImage(ctx context.Context, propertyId int, url string) error {
	query := "INSERT INTO property_images (url, property_id) VALUES (?, ?)"
	_, err := r.DB.ExecContext(ctx, query, url, propertyId)
	return err
}

func (r *ImageRepository) DeleteImage(ctx context.Context, imageId int) error {
	query := "DELETE FROM property_images WHERE id = ?"
	_, err := r.DB.ExecContext(ctx, query, imageId)
	return err
}

func (r *ImageRepository) GetImageByID(ctx context.Context, imageId int) (*domain.PropertyImage, error) {
	query := "SELECT * FROM property_images WHERE id = ?"

	i := domain.PropertyImage{}
	err := r.DB.QueryRowContext(ctx, query, imageId).Scan(&i.ID, &i.Url, &i.PropertyID)

	if err == sql.ErrNoRows {
		return nil, errors.New("image not found")
	}

	return &i, err
}
