package handler

import (
	"net/http"
	"strconv"

	"vp_backend/internal/service"

	"github.com/gin-gonic/gin"
)

// FavoriteHandler menangani HTTP request
// yang berkaitan dengan fitur properti favorit user.
type FavoriteHandler struct {
	FavoriteService *service.FavoriteService
}

// AddToFavorites menambahkan properti ke dalam
// daftar favorit user yang sedang login.
//
// Endpoint ini memerlukan autentikasi JWT.
//
// Path parameter:
// - id : ID properti
//
// Response:
// - 200 OK : properti berhasil ditambahkan ke favorit
// - 400 Bad Request : ID properti tidak valid
// - 401 Unauthorized : user belum login
// - 500 Internal Server Error : kesalahan internal
func (h *FavoriteHandler) AddToFavorites(c *gin.Context) {

	// Mengambil ID properti dari URL parameter
	propertyID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid property ID"},
		)
		return
	}

	// Mengambil user_id dari context (hasil JWT middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{"error": "unauthorized"},
		)
		return
	}

	// Menambahkan properti ke daftar favorit
	err = h.FavoriteService.AddFavorite(
		c.Request.Context(),
		userID.(int),
		propertyID,
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	// Response sukses
	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "property added to favorites successfully",
		},
	)
}

// RemoveFromFavorites menghapus properti dari
// daftar favorit user yang sedang login.
//
// Endpoint ini memerlukan autentikasi JWT.
//
// Path parameter:
// - id : ID properti
//
// Response:
// - 200 OK : properti berhasil dihapus dari favorit
// - 400 Bad Request : ID properti tidak valid
// - 401 Unauthorized : user belum login
// - 500 Internal Server Error : kesalahan internal
func (h *FavoriteHandler) RemoveFromFavorites(c *gin.Context) {

	// Mengambil ID properti dari URL parameter
	propertyID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid property ID"},
		)
		return
	}

	// Mengambil user_id dari context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{"error": "unauthorized"},
		)
		return
	}

	// Menghapus properti dari daftar favorit
	err = h.FavoriteService.RemoveFavorite(
		c.Request.Context(),
		userID.(int),
		propertyID,
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	// Response sukses
	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "property removed from favorites successfully",
		},
	)
}

// GetFavoriteProperties mengambil seluruh daftar
// properti favorit milik user yang sedang login.
//
// Endpoint ini memerlukan autentikasi JWT.
//
// Response:
// - 200 OK : daftar properti favorit
// - 401 Unauthorized : user belum login
// - 500 Internal Server Error : kesalahan internal
func (h *FavoriteHandler) GetFavoriteProperties(c *gin.Context) {

	// Mengambil user_id dari context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{"error": "unauthorized"},
		)
		return
	}

	// Mengambil seluruh properti favorit user
	properties, err := h.FavoriteService.GetAll(
		c.Request.Context(),
		userID.(int),
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	// Response sukses
	c.JSON(
		http.StatusOK,
		properties,
	)
}

