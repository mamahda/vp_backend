package handler

import (
	"net/http"
	"log"
	"vp_backend/internal/service"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/message"
)

type UserHandler struct {
	UserService *service.UserService
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	log.Println("Retrieved user_id from context:", userID)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	
	profile, err := h.UserService.GetUser(int(userID.(uint)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user profile"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": message.NewPrinter(message.MatchLanguage("en")).Sprintf("User profile retrieved successfully"),
		"data": gin.H{
			"id":       profile.ID,
			"username": profile.Username,
			"email":    profile.Email,
			"phone":    profile.Phone,
			"role_id":  profile.Role_ID,
		},
	})
}
