package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"vp_backend/internal/service"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
			return
		}

		userService := c.MustGet("user_service").(*service.UserService)
		user, err := userService.Get(c.Request.Context(), int(userID.(uint)))
		if err != nil || user == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			return
		}

		if !user.IsAdmin() {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "admin access required"})
			return
		}

		c.Next()
	}
}
