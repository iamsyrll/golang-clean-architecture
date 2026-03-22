package routes

import (
	"golang-clean-arch/internal/delivery/http/handler"
	"golang-clean-arch/internal/usecase"

	"github.com/gin-gonic/gin"
)

type RoutesConfig struct {
	Router *gin.Engine
	AuthUc usecase.AuthUsecase
}

func RegisterRoutes(cfg RoutesConfig) {
	handler.NewAuthHandler(cfg.Router, cfg.AuthUc)
}
