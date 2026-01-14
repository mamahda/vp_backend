package handler

import (
	"net/http"
	"vp_backend/internal/service"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/message"
)

type UserHandler struct {
	UserService *service.UserService
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	
	profile, err := h.UserService.Get(c.Request.Context(), int(userID.(uint)))
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

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := h.UserService.UpdateUser(c.Request.Context(), int(userID.(uint)), req.Username, req.Email, req.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message.NewPrinter(message.MatchLanguage("en")).Sprintf("User profile updated successfully"),
	})
}
