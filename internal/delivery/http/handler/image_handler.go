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

func (h *ImageHandler) GetAllImages(c *gin.Context) {
	propertyId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID properti tidak valid",
		})
		return
	}

	data, err := h.ImageService.GetAllPropertyImages(c.Request.Context(), propertyId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal mengambil data: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data gambar berhasil diambil", // Perbaikan pesan
		"data":    data,
	})
}

func (h *ImageHandler) UpdateCover(c *gin.Context){
	propertyId, err := strconv.Atoi(c.Param("id"))
	imageId, err := strconv.Atoi(c.Param("image_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Image ID tidak valid",
		})
		return
	}

	err = h.ImageService.UpdateCoverImage(c.Request.Context(), propertyId, imageId)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"success": false,
				"message": err.Error(),
			})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Cover Image berhasil diperbarui", // Perbaikan pesan
	})
}
