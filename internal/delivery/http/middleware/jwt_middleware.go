package middleware

import (
	"net/http"
	"strings"

	"vp_backend/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JwtClaims merepresentasikan payload JWT
// yang digunakan dalam aplikasi.
//
// UserID menyimpan ID user yang terautentikasi.
type JwtClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

// JWTAuth merupakan middleware autentikasi berbasis JWT.
//
// Middleware ini bertugas untuk:
// - Mengambil token dari header Authorization
// - Memvalidasi token JWT
// - Mengekstrak user_id dari claims
// - Menyimpan user_id ke dalam Gin context
//
// Format header yang diharapkan:
// Authorization: Bearer <jwt_token>
//
// Jika token tidak ada, tidak valid, atau gagal diverifikasi,
// request akan dihentikan dengan status Unauthorized (401).
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Mengambil header Authorization
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"success": false,
					"message": "missing token",
				},
			)
			return
		}

		// Menghapus prefix "Bearer " dari token
		tokenString := strings.Replace(auth, "Bearer ", "", 1)

		// Inisialisasi struktur claims JWT
		claims := &JwtClaims{}

		// Parsing dan validasi token JWT
		token, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			func(t *jwt.Token) (any, error) {
				return []byte(config.GetJWT()), nil
			},
		)

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"success": false,
					"message": "invalid token",
				},
			)
			return
		}

		// Menyimpan user_id ke dalam context
		// untuk digunakan oleh middleware atau handler selanjutnya
		c.Set("user_id", claims.UserID)

		// Melanjutkan request ke middleware / handler berikutnya
		c.Next()
	}
}
