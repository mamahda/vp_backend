package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"vp_backend/internal/usecase"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(user usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: user}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userUsecase.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
