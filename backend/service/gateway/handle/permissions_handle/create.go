package permissions_handle

import (
	"backend/protos/gen/go/sso/permissions"
	"backend/service/gateway/models/sso"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) create(c *gin.Context) {
	const op = "gateway.Permissions.Create"
	logger := h.logger.With(slog.String("op", op))

	var req sso_models.CreatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to bind JSON", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if req.Code == "" || req.Category == "" || req.AppID == 0 {
		logger.Warn("Missing required fields", slog.String("code", req.Code), slog.String("category", req.Category), slog.Int("app_id", int(req.AppID)))
		c.JSON(http.StatusBadRequest, gin.H{"error": "code, category and app_id are required"})
		return
	}

	grpcReq := &permissions.CreateRequest{
		Code:        req.Code,
		Description: req.Description,
		Category:    req.Category,
		AppId:       req.AppID,
	}

	resp, err := h.service.CreatePermission(c.Request.Context(), grpcReq)
	if err != nil {
		logger.Error("Failed to create permission", slog.String("error", err.Error()))
		h.handleGRPCError(c, err)
		return
	}

	permission, err := sso_models.PermissionFromProto(resp)
	if err != nil {
		logger.Error("Proto conversion failed", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	logger.Info("Permission created successfully", slog.String("permission_id", permission.ID), slog.Int("app_id", int(permission.AppID)))
	c.JSON(http.StatusCreated, permission)
}
