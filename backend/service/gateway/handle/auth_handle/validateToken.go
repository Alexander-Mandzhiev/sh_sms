package auth_handle

import (
	"backend/pkg/jwt_manager"
	"backend/protos/gen/go/auth"
	"backend/service/gateway/models/auth"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
	"time"
)

func (h *Handler) validateToken(c *gin.Context) {
	const op = "auth_handler.ValidateToken"
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
		logger.Warn("invalid authorization header", "error", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	grpcRes, err := h.service.ValidateToken(ctx, &auth.ValidateRequest{Token: token, TokenTypeHint: req.TokenTypeHint})
	if err != nil {
		logger.Error("token validation failed", slog.String("error", err.Error()), slog.String("token_type", req.TokenTypeHint), slog.String("token_hash", jwt_manager.HashToken(token)[:6]+"***"))
		h.handleGRPCError(c, err)
		return
	}

	validationResult, err := auth_models.TokenInfoFromProto(grpcRes)
	if err != nil || validationResult.ClientID == uuid.Nil {
		logger.Error("invalid token info response", slog.String("error", err.Error()), slog.String("token_hash", jwt_manager.HashToken(token)[:6]+"***"))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	logger.Info("token validated", slog.Bool("active", validationResult.Active), slog.String("client_id", validationResult.ClientID.String()), slog.String("user_id", validationResult.UserID.String()))
	c.JSON(http.StatusOK, validationResult)
}
