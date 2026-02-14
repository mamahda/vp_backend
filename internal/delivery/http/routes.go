package http

import (
	"github.com/gin-gonic/gin"

	"vp_backend/internal/delivery/http/handler"
	"vp_backend/internal/delivery/http/middleware"
)

// Handler berisi seluruh dependency handler HTTP
// yang akan digunakan dalam proses routing API.
type Handler struct {
	AuthHandler     *handler.AuthHandler
	UserHandler     *handler.UserHandler
	PropertyHandler *handler.PropertyHandler
	FavoriteHandler *handler.FavoriteHandler
	ImageHandler		*handler.ImageHandler
}

// RegisterRoutes mendaftarkan seluruh endpoint HTTP
// untuk aplikasi Victoria Property API.
func RegisterRoutes(r *gin.Engine, h Handler) {

	// Base API group dengan prefix /api
	api := r.Group("/api")

	// ==========================
	// AUTHENTICATION ROUTES
	// ==========================

	// Endpoint untuk registrasi user baru
	api.POST("/register", h.AuthHandler.Register)

	// Endpoint untuk login user dan menghasilkan JWT token
	api.POST("/login", h.AuthHandler.Login)

	// ==========================
	// PUBLIC PROPERTY ROUTES
	// ==========================

	// Endpoint untuk mengambil daftar properti
	// dengan dukungan query filter (price, location, dll)
	api.GET("/properties", h.PropertyHandler.GetProperties)

	// Endpoint untuk mengambil jumlah daftar properti
	// dengan dukungan query filter (price, location, dll)
	api.GET("/properties/count", h.PropertyHandler.GetCountData)

	// Endpoint untuk mengambil seluruh properti tanpa filter
	api.GET("/properties/all", h.PropertyHandler.GetAll)

	// Endpoint untuk mengambil detail properti berdasarkan ID
	api.GET("/properties/:id", h.PropertyHandler.GetByID)

	api.GET("/properties/:id/images", h.PropertyHandler.GetAllImages)

	// ==========================
	// PROTECTED ROUTES (JWT)
	// ==========================

	// Group route yang memerlukan autentikasi JWT
	protected := api
	protected.Use(middleware.JWTAuth())
	{

		// Endpoint untuk mengambil data profil user yang sedang login
		protected.GET("/profile", h.UserHandler.GetProfile)

		// Endpoint untuk memperbarui data profil user
		protected.PUT("/profile", h.UserHandler.UpdateProfile)

		// Endpoint untuk menambahkan properti ke daftar favorit user
		protected.POST("/properties/:id/favorite", h.FavoriteHandler.AddToFavorites)

		// Endpoint untuk menghapus properti dari daftar favorit user
		protected.DELETE("/properties/:id/favorite", h.FavoriteHandler.RemoveFromFavorites)

		// Endpoint untuk mengambil seluruh properti favorit milik user
		protected.GET("/favorites", h.FavoriteHandler.GetFavoriteProperties)

		// ==========================
		// AGENT / ADMIN ROUTES
		// ==========================

		// Group route khusus agent/admin
		protectedAgent := protected.Group("/agent")
		protectedAgent.Use(middleware.AdminAuth())
		{

			// Endpoint untuk menambahkan properti baru
			// (hanya dapat diakses oleh agent/admin)
			protectedAgent.POST("/properties", h.PropertyHandler.Create)

			protectedAgent.POST("/properties/:id/images", h.ImageHandler.UploadImages)

			// Endpoint untuk memperbarui data properti berdasarkan ID
			protectedAgent.PUT("/properties/:id", h.PropertyHandler.Update)

			// Endpoint untuk menghapus properti berdasarkan ID
			protectedAgent.DELETE("/properties/:id", h.PropertyHandler.Delete)
		}
	}
}
