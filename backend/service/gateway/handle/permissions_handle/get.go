package permissions_handle

import (
	"backend/service/gateway/models/sso"
	"log/slog"
	"net/http"
	"strconv"

	"backend/protos/gen/go/sso/permissions"
	"github.com/gin-gonic/gin"
)

func (h *Handler) get(c *gin.Context) {
	const op = "gateway.Permissions.Get"
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

	grpcReq := &permissions.GetRequest{
		Id:    id,
		AppId: int32(appID),
	}

	resp, err := h.service.GetPermission(c.Request.Context(), grpcReq)
	if err != nil {
		logger.Error("Failed to get permission", slog.String("permission_id", id), slog.Int("app_id", appID), slog.String("error", err.Error()))
		h.handleGRPCError(c, err)
		return
	}

	permission, err := sso_models.PermissionFromProto(resp)
	if err != nil {
		logger.Error("Proto conversion failed", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	logger.Info("Permission retrieved successfully", slog.String("permission_id", permission.ID), slog.Int("app_id", int(permission.AppID)))
	c.JSON(http.StatusOK, permission)
}
