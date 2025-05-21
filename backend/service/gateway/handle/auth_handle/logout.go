package auth_handle

import (
	"backend/pkg/cookies"
	"backend/pkg/jwt_manager"
	"backend/protos/gen/go/auth"
	"backend/service/gateway/models/auth"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"time"
)

func (h *Handler) logout(c *gin.Context) {
	const op = "auth_handler.Logout"
	logger := h.logger.With(slog.String("op", op))

	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		logger.Error("refresh token cookie missing", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "refresh token required"})
		return
	}

	var req auth_models.AuthLogoutRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		logger.Error("invalid request format", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if req.AccessToken == "" || refreshToken == "" {
		logger.Error("missing tokens in request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "access_token and refresh_token are required"})
		return
	}

	grpcReq := &auth.LogoutRequest{
		AccessToken:  req.AccessToken,
		RefreshToken: refreshToken,
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err = h.service.Logout(ctx, grpcReq); err != nil {
		logger.Error("logout failed",
			slog.String("error", err.Error()),
			slog.String("access_token", jwt_manager.HashToken(req.AccessToken)),
			slog.String("refresh_token", jwt_manager.HashToken(refreshToken)))
		h.handleGRPCError(c, err)
		return
	}

	cookies.RemoveRefreshCookie(c.Writer, cookies.DefaultConfig)

	logger.Info("successful logout",
		slog.String("access_token", jwt_manager.HashToken(req.AccessToken)),
		slog.String("refresh_token", jwt_manager.HashToken(refreshToken)))
	c.Status(http.StatusOK)
}
