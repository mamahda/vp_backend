package main

import (
	"fmt"
	"os"

	"vp_backend/internal/config"
	"vp_backend/internal/delivery/http"
	"vp_backend/internal/delivery/http/handler"
	"vp_backend/internal/repository"
	"vp_backend/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	db := config.InitDB()
	defer db.Close()

	r := gin.Default()

	userRepo := &repository.UserRepository{DB: db}
	authService := &service.AuthService{UserRepo: userRepo}
	authHandler := &handler.AuthHandler{AuthService: authService}

	userService := &service.UserService{UserRepo: userRepo}
	userHandler := &handler.UserHandler{UserService: userService}

	propertyRepo := &repository.PropertyRepository{DB: db}
	propertyService := &service.PropertyService{PropertyRepo: propertyRepo}
	propertyHandler := &handler.PropertyHandler{PropertyService: propertyService}

	favoriteRepo := &repository.FavoriteRepository{DB: db}
	favoriteService := &service.FavoriteService{FavoriteRepo: favoriteRepo}
	favoriteHandler := &handler.FavoriteHandler{FavoriteService: favoriteService}

	h := http.Handler{
		AuthHandler:     authHandler,
		UserHandler:     userHandler,
		PropertyHandler: propertyHandler,
		FavoriteHandler: favoriteHandler,
	}

	http.RegisterRoutes(r, h)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on port", port)
	r.Run(":" + port)
}
