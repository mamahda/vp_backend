package service

import (
	"context"
	"mime/multipart"
	"net/http"

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

func (s *PropertyService) AddPropertyImages(ctx context.Context, propertyId int, files []*multipart.FileHeader) error {
	// Definisi whitelist tipe file yang diizinkan
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/webp": true, // Opsional: format web modern
		"image/jpg":  true,
	}

	for i, file := range files {
		// --- SECURITY CHECK START ---

		// 1. Validasi Ukuran File (Opsional tapi disarankan, misal max 5MB)
		if file.Size > 5*1024*1024 {
			return domain.ErrFileTooLarge
		}

		// 2. Buka file untuk membaca header-nya
		src, err := file.Open()
		if err != nil {
			return err
		}

		// Baca 512 byte pertama (Magic Bytes) untuk deteksi konten
		buffer := make([]byte, 512)
		if _, err := src.Read(buffer); err != nil {
			src.Close() // Jangan lupa tutup jika error
			return err
		}

		// Deteksi MIME type yang sebenarnya
		contentType := http.DetectContentType(buffer)

		// Validasi apakah tipe file ada di whitelist
		if !allowedTypes[contentType] {
			src.Close()
			return domain.ErrInvalidFileType
		}

		// PENTING: Tutup file reader validasi ini agar resource tidak bocor.
		// Fungsi Storage.Upload nantinya akan membuka file ini lagi dari awal (fresh),
		// jadi kita tidak perlu melakukan Seek(0,0) di sini karena kita bekerja dengan FileHeader.
		src.Close()

		// --- SECURITY CHECK END ---

		// STEP A: Simpan file ke disk (panggil storage domain)
		webURL, err := s.Storage.Upload(file, "properties")
		if err != nil {
			return err
		}

		// STEP B: Jika ini adalah foto pertama (indeks 0), update tabel 'properties'
		if i == 0 {
			// Kita panggil repo untuk update kolom cover_image_url di tabel utama
			if err := s.PropertyRepo.UpdateCoverImage(ctx, propertyId, webURL); err != nil {
				return err
			}
		}

		// STEP C: Simpan SEMUA URL ke tabel galeri 'property_images' (untuk slider/detail)
		if err := s.PropertyRepo.SaveImage(ctx, propertyId, webURL); err != nil {
			return err
		}
	}
	return nil
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
