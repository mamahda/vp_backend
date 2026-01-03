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

	userRepo := &repository.UserRepository{DB: db}
	authService := &service.AuthService{UserRepo: userRepo}
	authHandler := &handler.AuthHandler{AuthService: authService}

	r := gin.Default()

	http.RegisterRoutes(r, authHandler)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on port", port)
	r.Run(":" + port)
}
