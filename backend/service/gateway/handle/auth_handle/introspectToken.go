package auth_handle

import (
	"backend/pkg/jwt_manager"
	"backend/protos/gen/go/auth"
	"backend/service/gateway/models/auth"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"time"
)

func (h *Handler) introspectToken(c *gin.Context) {
	const op = "auth_handler.IntrospectToken"
	logger := h.logger.With(slog.String("op", op))

	var req auth_models.TokenValidationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("invalid request format", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if req.TokenTypeHint != "access" && req.TokenTypeHint != "refresh" {
		logger.Warn("invalid token_type_hint", "hint", req.TokenTypeHint)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token_type_hint"})
		return
	}

	authHeader := c.GetHeader("Authorization")
	token, err := h.extractBearerToken(authHeader)
	if err != nil || token == "" {
		logger.Warn("invalid or empty token", "error", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	res, err := h.service.IntrospectToken(ctx, &auth.ValidateRequest{Token: token, TokenTypeHint: req.TokenTypeHint})
	if err != nil {
		logger.Error("introspection failed", slog.String("error", err.Error()), slog.String("token_hash", jwt_manager.HashToken(token)[:6]+"***"))
		h.handleGRPCError(c, err)
		return
	}

	introspection, err := auth_models.TokenIntrospectionFromProto(res)
	if err != nil {
		logger.Error("invalid introspection data", "error", err, "token_hash", jwt_manager.HashToken(token)[:6]+"***")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	logger.Info("token introspected", slog.Bool("active", introspection.Active), slog.String("client_id", introspection.ClientID.String()), slog.String("user_id", introspection.UserID.String()))
	c.JSON(http.StatusOK, introspection)
}
