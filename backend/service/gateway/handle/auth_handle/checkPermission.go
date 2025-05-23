package auth_handle

import (
	"backend/protos/gen/go/auth"
	"backend/service/gateway/models/auth"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
	"time"
)

func (h *Handler) checkPermission(c *gin.Context) {
	const op = "auth_handler.CheckPermission"
	logger := h.logger.With(slog.String("op", op))

	var req auth_models.CheckPermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("invalid request format", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if req.ClientID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "client_id is required"})
		return
	}
	if req.AppID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid app_id"})
		return
	}
	if req.Permission == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "permission is required"})
		return
	}

	token, err := h.extractTokenFromCookie(c, "access")
	if err != nil {
		logger.Warn("token extraction failed", "error", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token required"})
		return
	}

	grpcReq := &auth.PermissionCheckRequest{
		ClientId:   req.ClientID.String(),
		AppId:      int32(req.AppID),
		Permission: req.Permission,
		Resource:   req.Resource,
		Token:      token,
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	grpcRes, err := h.service.CheckPermission(ctx, grpcReq)
	if err != nil {
		logger.Error("permission check failed", slog.String("error", err.Error()), slog.String("client_id", req.ClientID.String()), slog.String("permission", req.Permission))
		h.handleGRPCError(c, err)
		return
	}

	response := auth_models.CheckPermissionFromProtoResponse(grpcRes)

	logger.Info("permission checked", slog.Bool("allowed", response.Allowed), slog.String("client_id", req.ClientID.String()), slog.String("permission", req.Permission))
	c.JSON(http.StatusOK, response)
}
