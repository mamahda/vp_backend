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
type ImageService struct {
	ImageRepo 	*repository.ImageRepository
	Storage      storage.Storage
}

func (s *ImageService) AddPropertyImages(ctx context.Context, propertyId int, files []*multipart.FileHeader) error {
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
			if err := s.ImageRepo.UpdateCoverImage(ctx, propertyId, webURL); err != nil {
				return err
			}
		}

		// STEP C: Simpan SEMUA URL ke tabel galeri 'property_images' (untuk slider/detail)
		if err := s.ImageRepo.SaveImage(ctx, propertyId, webURL); err != nil {
			return err
		}
	}
	return nil
}
