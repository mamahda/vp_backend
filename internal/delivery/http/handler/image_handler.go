package handler

import (
	"net/http"
	"strconv"

	"vp_backend/internal/service"

	"github.com/gin-gonic/gin"
)

type ImageHandler struct {
	ImageService *service.ImageService
}

func (h *ImageHandler) UploadImages(c *gin.Context) {
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
	err = h.ImageService.AddPropertyImages(c.Request.Context(), propertyId, files)
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

func (h *ImageHandler) RemoveImage(c *gin.Context) {
	imageId, _ := strconv.Atoi(c.Param("image_id"))
	propertyId, _ := strconv.Atoi(c.Param("image_id"))

	err := h.ImageService.RemovePropertyImage(c.Request.Context(), imageId, propertyId)
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
			"message": "Remove Berhasil",
		},
	)
}
