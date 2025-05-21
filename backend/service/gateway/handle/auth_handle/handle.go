package auth_handle

import (
	config "backend/pkg/config/gateway"
	"backend/protos/gen/go/auth"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type AuthService interface {
	Login(ctx context.Context, req *auth.LoginRequest) (*auth.AuthResponse, error)
	Logout(ctx context.Context, req *auth.LogoutRequest) error
	RefreshToken(ctx context.Context, req *auth.RefreshRequest) (*auth.AuthResponse, error)
	ValidateToken(ctx context.Context, req *auth.ValidateRequest) (*auth.TokenInfo, error)
	IntrospectToken(ctx context.Context, req *auth.ValidateRequest) (*auth.TokenIntrospection, error)
	CheckPermission(ctx context.Context, req *auth.PermissionCheckRequest) (*auth.PermissionCheckResponse, error)
	ListSessionsForUser(ctx context.Context, req *auth.SessionFilter) (*auth.SessionList, error)
	ListAllSessions(ctx context.Context, req *auth.AllSessionsFilter) (*auth.SessionList, error)
	TerminateSession(ctx context.Context, req *auth.SessionID) error
}

type Handler struct {
	logger  *slog.Logger
	service AuthService
	cfg     config.Config
}

func New(service AuthService, logger *slog.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger.With(slog.String("module", "auth_handler")),
	}
}

func (h *Handler) InitRoutes(router *gin.RouterGroup) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", h.login)
		authGroup.POST("/logout", h.logout)
		authGroup.POST("/refresh", h.refreshToken)
		authGroup.POST("/validate", h.validateToken)
		authGroup.POST("/introspect", h.introspectToken)
		authGroup.POST("/check-permission", h.checkPermission)
		authGroup.GET("/sessions", h.listUserSessions)
		authGroup.GET("/all-sessions", h.listAllSessions)
		authGroup.DELETE("/sessions/:session_id", h.terminateSession)
	}
}
