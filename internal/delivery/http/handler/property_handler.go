package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"vp_backend/internal/domain"
	"vp_backend/internal/service"
)

type PropertyHandler struct {
	PropertyService *service.PropertyService
}

func (h *PropertyHandler) Create(c *gin.Context) {
	var PropertyService domain.Property
	if err := c.ShouldBindJSON(&PropertyService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.PropertyService.Create(c.Request.Context(), &PropertyService); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "PropertyService created"})
}

func (h *PropertyHandler) GetAll(c *gin.Context) {
	properties, err := h.PropertyService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, properties)
}

func (h *PropertyHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	PropertyService, err := h.PropertyService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, PropertyService)
}

func (h *PropertyHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var PropertyService domain.Property
	if err := c.ShouldBindJSON(&PropertyService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	PropertyService.ID = id

	if err := h.PropertyService.Update(c.Request.Context(), &PropertyService); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "service updated"})
}

func (h *PropertyHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.PropertyService.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "service deleted"})
}

