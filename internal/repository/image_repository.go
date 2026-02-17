package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

func (r *ImageRepository) FindAllPropertyImages(ctx context.Context, propertyId int) ([]domain.PropertyImage, error) {
	query := "SELECT id, url, property_id FROM property_images WHERE property_id = ?"

	rows, err := r.DB.QueryContext(ctx, query, propertyId)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	// SANGAT PENTING: Selalu tutup rows untuk mengembalikan koneksi ke pool
	defer rows.Close()

	var images []domain.PropertyImage

	for rows.Next() {
		var i domain.PropertyImage
		// Sesuaikan urutan Scan dengan urutan kolom di SELECT
		if err := rows.Scan(&i.ID, &i.Url, &i.PropertyID); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		images = append(images, i)
	}

	// Cek apakah ada error saat iterasi rows
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return images, nil
}

func (r *ImageRepository) GetImageByID(ctx context.Context, imageId int, propertyId int) (*domain.PropertyImage, error) {
	query := "SELECT * FROM property_images WHERE id = ? AND property_id = ?"

	i := domain.PropertyImage{}
	err := r.DB.QueryRowContext(ctx, query, imageId, propertyId).Scan(&i.ID, &i.Url, &i.PropertyID)

	if err == sql.ErrNoRows {
		return nil, errors.New("image not found")
	}

	return &i, err
}
