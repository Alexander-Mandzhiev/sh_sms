package permissions_handle

import (
	"log/slog"
	"net/http"
	"strconv"

	"backend/protos/gen/go/sso/permissions"
	"backend/service/gateway/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) update(c *gin.Context) {
	const op = "gateway.Permissions.Update"
	logger := h.logger.With(slog.String("op", op))

	id := c.Param("id")
	appIDStr := c.Query("app_id")

	if id == "" || appIDStr == "" {
		logger.Error("Missing required parameters", slog.String("permission_id", id), slog.String("app_id", appIDStr))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters 'id' and 'app_id' are required"})
		return
	}

	appID, err := strconv.Atoi(appIDStr)
	if err != nil {
		logger.Error("Invalid app_id format", slog.String("app_id", appIDStr), slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid app_id format, must be integer"})
		return
	}

	var req models.UpdatePermissionRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to bind JSON", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if req.Code == nil && req.Description == nil &&
		req.Category == nil && req.IsActive == nil {
		logger.Warn("No fields provided for update")
		c.JSON(http.StatusBadRequest, gin.H{"error": "At least one field must be provided for update"})
		return
	}

	grpcReq := &permissions.UpdateRequest{
		Id:          id,
		AppId:       int32(appID),
		Code:        req.Code,
		Description: req.Description,
		Category:    req.Category,
		IsActive:    req.IsActive,
	}

	resp, err := h.service.UpdatePermission(c.Request.Context(), grpcReq)
	if err != nil {
		logger.Error("Failed to update permission", slog.String("permission_id", id), slog.Int("app_id", appID), slog.String("error", err.Error()))
		h.handleGRPCError(c, err)
		return
	}

	permission, err := models.PermissionFromProto(resp)
	if err != nil {
		logger.Error("Proto conversion failed", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	logger.Info("Permission updated successfully", slog.String("permission_id", permission.ID), slog.Int("app_id", int(permission.AppID)))
	c.JSON(http.StatusOK, permission)
}
