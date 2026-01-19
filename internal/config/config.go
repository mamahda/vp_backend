package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv memuat environment variable dari file .env ke dalam
// sistem environment aplikasi.
//
// Jika file .env tidak ditemukan, aplikasi tetap berjalan
// dengan menggunakan environment variable dari sistem.
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
		return
	}

	log.Println(".env file loaded successfully")
}

// GetEnv mengambil nilai environment variable berdasarkan key.
// Jika environment variable tidak ditemukan atau kosong,
// maka nilai defaultValue akan digunakan.
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// GetJWT mengambil secret key untuk JWT dari environment variable
// JWT_SECRET.
//
// Jika JWT_SECRET tidak diset, maka akan menggunakan
// nilai default sebagai fallback (digunakan untuk development).
func GetJWT() string {
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		return "yoursecretkey"
	}
	return key
}

