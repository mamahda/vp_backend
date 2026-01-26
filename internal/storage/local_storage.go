package storage

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

// internal/storage/local_storage.go
type localStorage struct {
	publicPath string // "./public/uploads"
	baseUrl    string // "/static"
}

func NewLocalStorage(publicPath, baseUrl string) Storage {
	return &localStorage{
		publicPath: publicPath,
		baseUrl:    baseUrl,
	}
}

func (l *localStorage) Upload(file *multipart.FileHeader, folder string) (string, error) {
	// 1. Buat Nama Unik
	ext := filepath.Ext(file.Filename)
	fileName := uuid.New().String() + ext

	// 2. Tentukan Path Simpan
	dst := filepath.Join(l.publicPath, folder, fileName)

	// 3. Pakai fitur Gin yang "Quality of Life" itu
	if err := saveFile(file, dst); err != nil {
		return "", err
	}

	// 4. Balikkan URL untuk disimpan di DB nanti
	return fmt.Sprintf("%s/%s/%s", l.baseUrl, folder, fileName), nil
}

func saveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
