package handler

import (
	"net/http"
	"strconv"

	"vp_backend/internal/service"

	"github.com/gin-gonic/gin"
)

type FavoriteHandler struct {
	FavoriteService *service.FavoriteService
}

func (h *FavoriteHandler) AddToFavorites(c *gin.Context) {
	propertyID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid property ID"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err = h.FavoriteService.AddFavorite(c.Request.Context(), userID.(int), propertyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Property added to favorites successfully",
	})
}

func (h *FavoriteHandler) RemoveFromFavorites(c *gin.Context) {
	propertyID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid property ID"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err = h.FavoriteService.RemoveFavorite(c.Request.Context(), userID.(int), propertyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Property remove from favorites successfully",
	})
}

func (h *FavoriteHandler) GetFavoriteProperties(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	properties, err := h.FavoriteService.GetAll(c.Request.Context(), userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, properties)
}
