package main

import (
	"fmt"
	"os"

	"vp_backend/internal/config"
	"vp_backend/internal/delivery/http"
	"vp_backend/internal/delivery/http/handler"
	"vp_backend/internal/repository"
	"vp_backend/internal/service"
	"vp_backend/internal/storage"

	"github.com/gin-gonic/gin"
)

// main merupakan entry point aplikasi backend Victoria Property.
// Fungsi ini bertanggung jawab untuk:
// - Memuat konfigurasi environment
// - Menginisialisasi database
// - Menyusun dependency (repository, service, handler)
// - Mendaftarkan HTTP routes
// - Menjalankan HTTP server
func main() {

	// Memuat variabel environment dari file .env
	config.LoadEnv()

	// Inisialisasi koneksi database
	db := config.InitDB()
	defer db.Close()

	// Inisialisasi Gin HTTP server dengan default middleware
	// (logger dan recovery)
	r := gin.Default()

	// ==========================
	// STATIC FILES (GLOBAL)
	// ==========================
	r.Static("/static", "./public/uploads")

	propertyStorage := storage.NewLocalStorage("./public/uploads", "/static")

	// ==========================
	// REPOSITORY INITIALIZATION
	// ==========================

	// Repository bertanggung jawab untuk interaksi langsung
	// dengan database
	userRepo := &repository.UserRepository{DB: db}
	propertyRepo := &repository.PropertyRepository{DB: db}
	favoriteRepo := &repository.FavoriteRepository{DB: db}

	// ==========================
	// SERVICE INITIALIZATION
	// ==========================

	// Service berisi business logic aplikasi
	authService := &service.AuthService{UserRepo: userRepo}
	userService := &service.UserService{UserRepo: userRepo}
	propertyService := &service.PropertyService{PropertyRepo: propertyRepo, Storage: propertyStorage}
	favoriteService := &service.FavoriteService{FavoriteRepo: favoriteRepo}

	// ==========================
	// HANDLER INITIALIZATION
	// ==========================

	// Handler berfungsi sebagai layer HTTP
	// yang menangani request dan response
	authHandler := &handler.AuthHandler{AuthService: authService}
	userHandler := &handler.UserHandler{UserService: userService}
	propertyHandler := &handler.PropertyHandler{PropertyService: propertyService}
	favoriteHandler := &handler.FavoriteHandler{FavoriteService: favoriteService}

	// ==========================
	// INJECT SERVICES TO CONTEXT
	// ==========================
	r.Use(func(c *gin.Context) {
		// Set user_service ke context agar bisa diakses middleware
		c.Set("user_service", userService)
		c.Next()
	})

	// ==========================
	// ROUTE REGISTRATION
	// ==========================

	// Menggabungkan seluruh handler ke dalam
	// satu struct untuk kebutuhan routing
	h := http.Handler{
		AuthHandler:     authHandler,
		UserHandler:     userHandler,
		PropertyHandler: propertyHandler,
		FavoriteHandler: favoriteHandler,
	}

	// Mendaftarkan seluruh endpoint API
	http.RegisterRoutes(r, h)

	// ==========================
	// SERVER STARTUP
	// ==========================

	// Mengambil port dari environment variable
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on port", port)

	// Menjalankan HTTP server
	r.Run(":" + port)
}
