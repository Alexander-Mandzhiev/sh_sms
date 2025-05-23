// refresh.go
package auth_handle

import (
	"backend/pkg/cookies"
	"backend/pkg/utils"
	"backend/protos/gen/go/auth"
	"backend/service/gateway/models/auth"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"time"
)

func (h *Handler) refreshToken(c *gin.Context) {
	const op = "auth_handler.RefreshToken"
	logger := h.logger.With(slog.String("op", op))

	logger.Debug("Request cookies",
		slog.Any("cookies", c.Request.Cookies()),
		slog.String("host", c.Request.Host),
		slog.String("url", c.Request.URL.String()),
	)

	refreshCookie, err := c.Cookie(refreshToken)
	if errors.Is(err, http.ErrNoCookie) {
		logger.Error("Cookie 'refresh_token' not found in request")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}
	if err != nil {
		logger.Error("Error reading cookie", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	var req auth_models.RefreshRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		logger.Error("invalid request format", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err = utils.ValidateUUID(req.ClientID); err != nil {
		logger.Error("invalid client_id", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid client_id format"})
		return
	}

	if err = utils.ValidateAppID(req.AppID); err != nil {
		logger.Error("invalid app_id", "app_id", req.AppID)
		c.JSON(http.StatusBadRequest, gin.H{"error": "app_id must be positive"})
		return
	}

	grpcReq := &auth.RefreshRequest{
		RefreshToken: refreshCookie,
		ClientId:     req.ClientID.String(),
		AppId:        int32(req.AppID),
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	res, err := h.service.RefreshToken(ctx, grpcReq)
	if err != nil {
		logger.Error("refresh failed", "error", err)
		h.handleGRPCError(c, err)
		return
	}

	cookies.SetRefreshCookie(c.Writer, res.AccessToken, accessToken, cookies.DefaultConfig, h.cfg.TokensTTL.AccessTokenDuration)
	cookies.SetRefreshCookie(c.Writer, res.RefreshToken, refreshToken, cookies.DefaultConfig, h.cfg.TokensTTL.RefreshTokenDuration)

	response, err := auth_models.AuthResponseFromProto(res)
	if err != nil {
		logger.Error("failed to convert response", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	logger.Info("successful refresh", "user_id", response.User.ID)
	c.JSON(http.StatusOK, response)
}
