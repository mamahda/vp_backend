package handler

import (
	"net/http"
	"strconv"

	"vp_backend/internal/domain"
	"vp_backend/internal/service"

	"github.com/gin-gonic/gin"
)

// PropertyHandler menangani seluruh HTTP request
// yang berkaitan dengan manajemen dan pengambilan data properti.
type PropertyHandler struct {
	PropertyService *service.PropertyService
}

// Create menangani proses pembuatan properti baru.
//
// Endpoint ini hanya dapat diakses oleh user
// dengan role admin atau agent.
//
// Request body:
// - JSON object Property
//
// Response:
// - 201 Created : properti berhasil dibuat
// - 400 Bad Request : request tidak valid
// - 500 Internal Server Error : kesalahan internal
func (h *PropertyHandler) Create(c *gin.Context) {

	// Binding request JSON ke domain Property
	var property domain.Property
	if err := c.ShouldBindJSON(&property); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	// Memproses pembuatan properti melalui service
	if err := h.PropertyService.Create(
		c.Request.Context(),
		&property,
	); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	// Response sukses
	c.JSON(
		http.StatusCreated,
		gin.H{"message": "property created successfully"},
	)
}

// GetProperties mengambil daftar properti
// berdasarkan filter dan pagination.
//
// Filter dapat dikirim melalui query parameter
// (misalnya harga, lokasi, tipe, dll).
//
// Response:
// - 200 OK : daftar properti + metadata pagination
// - 400 Bad Request : format parameter tidak valid
// - 500 Internal Server Error : kesalahan internal
func (h *PropertyHandler) GetProperties(c *gin.Context) {

	// Binding query parameter ke struct filter
	var filters domain.PropertyFilters
	if err := c.ShouldBindQuery(&filters); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error":   "invalid query parameter format",
				"details": err.Error(),
			},
		)
		return
	}

	// Mengambil properti berdasarkan filter
	data, total, err := h.PropertyService.GetFilteredProperty(
		c.Request.Context(),
		&filters,
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	// Response sukses dengan metadata pagination
	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"data":    data,
			"meta": gin.H{
				"total_count":  total,
				"current_page": filters.Page,
				"limit":        filters.Limit,
				"total_pages":  (total + filters.Limit - 1) / filters.Limit,
			},
		},
	)
}

// GetAll mengambil seluruh data properti
// tanpa filter maupun pagination.
//
// Response:
// - 200 OK : seluruh data properti
// - 500 Internal Server Error : kesalahan internal
func (h *PropertyHandler) GetAll(c *gin.Context) {

	properties, err := h.PropertyService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		properties,
	)
}

// GetByID mengambil detail properti
// berdasarkan ID.
//
// Path parameter:
// - id : ID properti
//
// Response:
// - 200 OK : detail properti
// - 404 Not Found : properti tidak ditemukan
func (h *PropertyHandler) GetByID(c *gin.Context) {

	// Mengambil ID properti dari URL parameter
	id, _ := strconv.Atoi(c.Param("id"))

	property, err := h.PropertyService.GetByID(
		c.Request.Context(),
		id,
	)
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		property,
	)
}

// Update menangani proses pembaruan data properti
// berdasarkan ID.
//
// Endpoint ini hanya dapat diakses oleh admin/agent.
//
// Path parameter:
// - id : ID properti
//
// Request body:
// - JSON object Property
//
// Response:
// - 200 OK : properti berhasil diperbarui
// - 400 Bad Request : request tidak valid
// - 500 Internal Server Error : kesalahan internal
func (h *PropertyHandler) Update(c *gin.Context) {

	// Mengambil ID properti dari URL parameter
	id, _ := strconv.Atoi(c.Param("id"))

	// Binding request JSON ke domain Property
	var property domain.Property
	if err := c.ShouldBindJSON(&property); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	// Set ID properti dari path parameter
	property.ID = id

	// Memproses update properti melalui service
	if err := h.PropertyService.Update(
		c.Request.Context(),
		&property,
	); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{"message": "property updated successfully"},
	)
}

// Delete menghapus data properti
// berdasarkan ID.
//
// Endpoint ini hanya dapat diakses oleh admin/agent.
//
// Path parameter:
// - id : ID properti
//
// Response:
// - 200 OK : properti berhasil dihapus
// - 500 Internal Server Error : kesalahan internal
func (h *PropertyHandler) Delete(c *gin.Context) {

	// Mengambil ID properti dari URL parameter
	id, _ := strconv.Atoi(c.Param("id"))

	// Menghapus properti melalui service
	if err := h.PropertyService.Delete(
		c.Request.Context(),
		id,
	); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{"message": "property deleted successfully"},
	)
}

