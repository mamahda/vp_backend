package handler

import (
	"net/http"

	"vp_backend/internal/service"

	"github.com/gin-gonic/gin"
)

// UserHandler menangani HTTP request
// yang berkaitan dengan data dan profil user.
type UserHandler struct {
	UserService *service.UserService
}

// GetProfile mengambil data profil user
// yang sedang login berdasarkan JWT.
//
// Endpoint ini memerlukan autentikasi JWT.
//
// Response:
// - 200 OK : data profil user berhasil diambil
// - 401 Unauthorized : user belum terautentikasi
// - 500 Internal Server Error : gagal mengambil data user
func (h *UserHandler) GetProfile(c *gin.Context) {

	// Mengambil user_id dari context (hasil JWT middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"success": false,
				"message": "unauthorized",
			},
		)
		return
	}

	// Mengambil data profil user melalui service
	profile, err := h.UserService.Get(
		c.Request.Context(),
		int(userID.(uint)),
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"success": false,
				"message": "failed to get user profile",
			},
		)
		return
	}

	// Response sukses
	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"message": "User profile retrieved successfully",
			"data": gin.H{
				"id":       profile.ID,
				"username": profile.Username,
				"email":    profile.Email,
				"phone":    profile.Phone,
				"role_id":  profile.Role_ID,
			},
		},
	)
}

// UpdateProfile memperbarui data profil user
// yang sedang login.
//
// Endpoint ini memerlukan autentikasi JWT.
//
// Request body:
// - username : nama pengguna
// - email    : email pengguna
// - phone    : nomor telepon pengguna
//
// Response:
// - 200 OK : profil user berhasil diperbarui
// - 400 Bad Request : request tidak valid
// - 401 Unauthorized : user belum terautentikasi
// - 500 Internal Server Error : gagal memperbarui data user
func (h *UserHandler) UpdateProfile(c *gin.Context) {

	// Mengambil user_id dari context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"success": false,
				"message": "unauthorized",
			},
		)
		return
	}

	// Struktur request update profil
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
	}

	// Binding request JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"message": "invalid request",
			},
		)
		return
	}

	// Memproses update profil melalui service
	err := h.UserService.UpdateUser(
		c.Request.Context(),
		int(userID.(uint)),
		req.Username,
		req.Email,
		req.Phone,
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"success": false,
				"message": "failed to update user profile",
			},
		)
		return
	}

	// Response sukses
	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"message": "User profile updated successfully",
		},
	)
}
