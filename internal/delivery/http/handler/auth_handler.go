package handler

import (
	"errors"
	"net/http"

	"vp_backend/internal/domain"
	"vp_backend/internal/service"

	"github.com/gin-gonic/gin"
)

// AuthHandler bertanggung jawab untuk menangani
// HTTP request terkait autentikasi user,
// seperti registrasi dan login.
type AuthHandler struct {
	AuthService *service.AuthService
}

// Register menangani proses registrasi user baru.
//
// Endpoint ini akan:
// - Menerima request body berupa data user
// - Melakukan validasi dan binding JSON
// - Meneruskan proses registrasi ke AuthService
//
// Response:
// - 201 Created : registrasi berhasil
// - 400 Bad Request : request tidak valid
// - 409 Conflict : email sudah terdaftar
// - 500 Internal Server Error : kesalahan internal
func (h *AuthHandler) Register(c *gin.Context) {

	// Binding request JSON ke domain User
	var req domain.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"message": err.Error(),
			},
		)
		return
	}

	// Memproses registrasi melalui service
	if err := h.AuthService.Register(
		c.Request.Context(),
		&req,
	); err != nil {

		// Menangani error email sudah terdaftar
		if errors.Is(err, domain.ErrEmailAlreadyExists) {
			c.JSON(
				http.StatusConflict,
				gin.H{
					"success": false,
					"message": err.Error(),
				},
			)
			return
		}

		// Error lainnya
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"success": false,
				"message": err.Error()},
		)
		return
	}

	// Response sukses
	c.JSON(
		http.StatusCreated,
		gin.H{
			"success": true,
			"message": "register success",
		},
	)
}

// Login menangani proses autentikasi user.
//
// Endpoint ini akan:
// - Menerima email dan password dari request body
// - Memvalidasi kredensial melalui AuthService
// - Menghasilkan JWT token jika berhasil
//
// Response:
// - 200 OK : login berhasil
// - 400 Bad Request : request tidak valid
// - 401 Unauthorized : email atau password salah
func (h *AuthHandler) Login(c *gin.Context) {

	// Struktur request login
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Binding request JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"message": err.Error()},
		)
		return
	}

	// Proses login melalui service
	user, token, err := h.AuthService.Login(
		c.Request.Context(),
		req.Email,
		req.Password,
	)
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"success": false,
				"message": err.Error()},
		)
		return
	}

	// Response sukses
	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"message": "login success",
			"data": gin.H{
				"token": token,
				"user":  user,
			},
		},
	)
}
