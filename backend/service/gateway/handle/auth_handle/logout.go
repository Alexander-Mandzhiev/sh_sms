package auth_handle

import (
	"backend/pkg/cookies"
	"backend/pkg/jwt_manager"
	"backend/protos/gen/go/auth"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"time"
)

func (h *Handler) logout(c *gin.Context) {
	const op = "auth_handler.Logout"
	logger := h.logger.With(slog.String("op", op))

	refreshCookie, err := c.Cookie(refreshToken)
	if errors.Is(err, http.ErrNoCookie) {
		logger.Error("Cookie 'refresh_token' not found")
	} else if err != nil {
		logger.Error("Error reading cookie", "error", err)
	}

	if refreshCookie == "" {
		logger.Error("missing tokens in request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "access_token and refresh_token are required"})
		return
	}

	grpcReq := &auth.LogoutRequest{
		RefreshToken: refreshCookie,
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err = h.service.Logout(ctx, grpcReq); err != nil {
		logger.Error("logout failed", slog.String("error", err.Error()), slog.String("refresh_token", jwt_manager.HashToken(refreshCookie)))
		h.handleGRPCError(c, err)
		return
	}

	cookies.RemoveRefreshCookie(c.Writer, accessToken, cookies.DefaultConfig)
	cookies.RemoveRefreshCookie(c.Writer, refreshToken, cookies.DefaultConfig)

	logger.Info("successful logout", slog.String("refresh_token", jwt_manager.HashToken(refreshCookie)))
	c.Status(http.StatusOK)
}
