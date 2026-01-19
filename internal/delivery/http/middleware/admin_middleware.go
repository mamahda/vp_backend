package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"vp_backend/internal/service"
)

// AdminAuth merupakan middleware untuk membatasi akses endpoint
// hanya kepada user dengan role admin atau agent.
//
// Middleware ini mengasumsikan bahwa:
// - Middleware autentikasi JWT telah dijalankan sebelumnya
// - user_id dan user_service sudah tersedia di Gin context
//
// Jika user tidak terautentikasi, tidak ditemukan,
// atau bukan admin, request akan dihentikan.
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Mengambil user_id dari context (hasil dari JWT middleware)
		userID, exists := c.Get("user_id")
		if !exists {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "user not authenticated"},
			)
			return
		}

		// Mengambil UserService dari context
		userService := c.MustGet("user_service").(*service.UserService)

		// Mengambil data user berdasarkan user_id
		user, err := userService.Get(
			c.Request.Context(),
			int(userID.(uint)),
		)
		if err != nil || user == nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "user not found"},
			)
			return
		}

		// Mengecek apakah user memiliki hak akses admin
		if !user.IsAdmin() {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				gin.H{"error": "admin access required"},
			)
			return
		}

		// Melanjutkan request jika user adalah admin
		c.Next()
	}
}

