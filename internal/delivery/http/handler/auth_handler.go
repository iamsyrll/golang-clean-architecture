package handler

import (
	"net/http"

	"golang-clean-arch/internal/entity"
	"golang-clean-arch/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	uc usecase.AuthUsecase
}

func NewAuthHandler(r *gin.Engine, uc usecase.AuthUsecase) {
	h := &AuthHandler{uc: uc}

	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var params entity.LoginParams

	err := c.ShouldBindBodyWithJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	token, err := h.uc.Login(params.Email, params.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "login success", "token": token})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var params entity.RegisterParams

	err := c.ShouldBindBodyWithJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
	}

	err = h.uc.Register(params.Username, params.Email, params.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "registered"})
}
