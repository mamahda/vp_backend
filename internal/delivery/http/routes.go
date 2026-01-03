package http

import (
	"github.com/gin-gonic/gin"

	"vp_backend/internal/delivery/http/handler"
	"vp_backend/internal/delivery/http/middleware"
)

type Handler struct {
	AuthHandler *handler.AuthHandler
}

func RegisterRoutes(r *gin.Engine, h Handler) {
	r.POST("/register", h.AuthHandler.Register)
	r.POST("/login", h.AuthHandler.Login)

	protected := r.Group("/api")
	protected.Use(middleware.JWTAuth())
	{
		protected.GET("/profile", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "authorized"})
		})
	}
}

