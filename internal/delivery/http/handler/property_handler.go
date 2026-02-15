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
			gin.H{
				"success": false,
				"message": err.Error(),
			},
		)
		return
	}

	userID := c.GetInt("user_id")
	property.AgentId = userID

	// Memproses pembuatan properti melalui service
	if err := h.PropertyService.Create(
		c.Request.Context(),
		&property,
	); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"success": false,
				"message": err.Error(),
			},
		)
		return
	}

	// Response sukses
	c.JSON(
		http.StatusCreated,
		gin.H{
			"success": true,
			"message": "property created successfully",
		},
	)
}

func (h *PropertyHandler) UploadImages(c *gin.Context) {
	propertyId, _ := strconv.Atoi(c.Param("id"))

	// 1. Ambil file dari request
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"message": "Invalid form",
			},
		)
		return
	}
	files := form.File["images"]

	// 2. Panggil Service (Passing context standar Go)
	err = h.PropertyService.AddPropertyImages(c.Request.Context(), propertyId, files)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"success": false,
				"message": err.Error(),
			})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"message": "Upload Berhasil",
		},
	)
}

func (h *PropertyHandler) GetCountData(c *gin.Context) {
	// Binding query parameter ke struct filter
	filters := new(domain.PropertyFilters)
	if err := c.ShouldBindQuery(filters); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"message": "invalid query parameter format",
			},
		)
		return
	}

	// Menghitung properti berdasarkan filter
	total, err := h.PropertyService.GetCountData(
		c.Request.Context(),
		filters,
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"success": false,
				"message": err.Error(),
			},
		)
		return
	}

	// Response sukses
	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"message": "get number of available property",
			"data": gin.H{
				"count": total,
			},
		},
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
				"success": false,
				"message": "invalid query parameter format",
			},
		)
		return
	}

	// Mengambil properti berdasarkan filter
	data, err := h.PropertyService.GetFilteredProperty(
		c.Request.Context(),
		&filters,
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"success": false,
				"message": err.Error(),
			},
		)
		return
	}

	// Response sukses dengan metadata pagination
	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"message": "get property data with filter and pagination",
			"data":    data,
			"meta": gin.H{
				"current_page": filters.Page,
				"limit":        filters.Limit,
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
			gin.H{
				"success": false,
				"message": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"message": "data all available property",
			"data": gin.H{
				"property": properties,
			},
		},
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
			gin.H{
				"success": false,
				"message": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"message": "get property by ID",
			"data": gin.H{
				"property": property,
			},
		},
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
			gin.H{
				"success": false,
				"message": err.Error(),
			},
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
			gin.H{
				"success": false,
				"message": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"message": "property updated successfully",
		},
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
			gin.H{
				"success": false,
				"message": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"message": "property deleted successfully",
		},
	)
}
