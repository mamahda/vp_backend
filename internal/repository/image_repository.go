package repository

import (
	"context"
	"database/sql"

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

func (r *ImageRepository) SaveImage(ctx context.Context, propertyID int, url string) error {
	query := "INSERT INTO property_images (url, property_id) VALUES (?, ?)"
	_, err := r.DB.ExecContext(ctx, query, url, propertyID)
	return err
}
