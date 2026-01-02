package main

import (
	"fmt"
	"os"

	"vp_backend/internal/config"
	"vp_backend/internal/delivery/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load ENV
	config.LoadEnv()

	// DB
	db := config.InitDB()
	defer db.Close()

	// Gin
	r := gin.Default()

	http.RegisterRoutes(r, db)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on port", port)
	r.Run(":" + port)
}
