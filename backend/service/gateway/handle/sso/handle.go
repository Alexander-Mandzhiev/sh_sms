package sso_handle

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

type SSOService interface {
}

type Handler struct {
	logger  *slog.Logger
	service SSOService
}

func New(service SSOService, logger *slog.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) InitRoutes(router *gin.RouterGroup) {
	users := router.Group("/users")
	{
		users.POST("", h.createUser)
		users.GET("/:id", h.getUser)
		users.PUT("/:id", h.updateUser)
		users.DELETE("/:id", h.deleteUser)
		users.GET("", h.listUsers)
		users.POST("/:id/password", h.setUserPassword)
		users.POST("/:id/restore", h.restoreUser)
	}
}
