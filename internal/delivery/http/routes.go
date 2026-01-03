package http

import (
	"github.com/gin-gonic/gin"

	"vp_backend/internal/delivery/http/handler"
	"vp_backend/internal/delivery/http/middleware"
)

func RegisterRoutes(r *gin.Engine, auth *handler.AuthHandler) {
	r.POST("/register", auth.Register)
	r.POST("/login", auth.Login)

	protected := r.Group("/api")
	protected.Use(middleware.JWTAuth())
	{
		protected.GET("/profile", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "authorized"})
		})
	}
}

