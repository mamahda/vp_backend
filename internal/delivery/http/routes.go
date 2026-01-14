package http

import (
	"github.com/gin-gonic/gin"

	"vp_backend/internal/delivery/http/handler"
	"vp_backend/internal/delivery/http/middleware"
)

type Handler struct {
	AuthHandler     *handler.AuthHandler
	UserHandler     *handler.UserHandler
	PropertyHandler *handler.PropertyHandler
	FavoriteHandler *handler.FavoriteHandler
}

func RegisterRoutes(r *gin.Engine, h Handler) {
	api := r.Group("/api")
	api.POST("/register", h.AuthHandler.Register)
	api.POST("/login", h.AuthHandler.Login)
	api.GET("/properties", h.PropertyHandler.GetAll)
	api.GET("/properties/:id", h.PropertyHandler.GetByID)

	protected := api
	protected.Use(middleware.JWTAuth())
	{
		protected.GET("/profile", h.UserHandler.GetProfile)
		protected.PUT("/profile", h.UserHandler.UpdateProfile)

		protected.POST("/properties/:id/favorite", h.FavoriteHandler.AddToFavorites)
		protected.DELETE("/properties/:id/favorite", h.FavoriteHandler.RemoveFromFavorites)
		protected.GET("/favorites", h.FavoriteHandler.GetFavoriteProperties)

		protectedAdmin := protected.Group("/agent")
		protectedAdmin.Use(middleware.AdminAuth())
		{
			protectedAdmin.POST("/properties", h.PropertyHandler.Create)
			protectedAdmin.PUT("/properties/:id", h.PropertyHandler.Update)
			protectedAdmin.DELETE("/properties/:id", h.PropertyHandler.Delete)
		}
	}
}
