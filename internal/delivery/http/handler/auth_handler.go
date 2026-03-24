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

type AuthResponseLogin struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type AuthResponseFailed struct {
	Error string `json:"error"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

func NewAuthHandler(r *gin.Engine, uc usecase.AuthUsecase) {
	h := &AuthHandler{uc: uc}

	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
}

// Login godoc
//
//	@Summary		Login User
//	@Description	Endpoint for user login
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//
// @Param request body entity.LoginParams true "Login request"
//
//	@Success		200	{object}	AuthResponseLogin
//	@Failure		400	{object}	AuthResponseFailed
//	@Failure		404	{object}	AuthResponseFailed
//	@Failure		500	{object}	AuthResponseFailed
//	@Router			/login [post]
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

// Register godoc
//
//	@Summary		Register User
//	@Description	Endpoint for user register
//	@Tags			Auth
//	@Accept			json
//
// @Security ApiKeyAuth
//
// @Param request body entity.RegisterParams true "Register request"
//
//	@Produce		json
//	@Success		200	{object}	RegisterResponse
//	@Failure		400	{object}	AuthResponseFailed
//	@Failure		404	{object}	AuthResponseFailed
//	@Failure		500	{object}	AuthResponseFailed
//	@Router			/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var params entity.RegisterParams

	err := c.ShouldBindBodyWithJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	err = h.uc.Register(params.Username, params.Email, params.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "registered"})
}
