package http

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"vp_backend/internal/delivery/http/handler"
	"vp_backend/internal/repository"
	"vp_backend/internal/usecase"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	// dependency injection
	userRepo := repository.NewUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	// routes
	router.GET("/users", userHandler.GetUsers)
}
