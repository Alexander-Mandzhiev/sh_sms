package role_permissions_handle

import (
	"log/slog"
	"net/http"

	"backend/protos/gen/go/sso/role_permissions"
	"backend/service/gateway/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) addPermissionsToRole(c *gin.Context) {
	const op = "gateway.RolePermissions.AddPermissionsToRole"
	logger := h.logger.With(slog.String("op", op))

	var req models.RolePermissionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("failed to bind JSON", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if req.ClientID == "" || req.AppID == 0 || len(req.PermissionIDs) == 0 {
		logger.Warn("missing required fields", slog.String("client_id", req.ClientID), slog.Int("app_id", int(req.AppID)), slog.Int("permissions_count", len(req.PermissionIDs)))
		c.JSON(http.StatusBadRequest, gin.H{"error": "client_id, app_id and permission_ids are required"})
		return
	}

	grpcReq := &role_permissions.PermissionsRequest{
		RoleId:        req.RoleId.String(),
		ClientId:      req.ClientID,
		AppId:         req.AppID,
		PermissionIds: req.PermissionIDs,
	}

	resp, err := h.service.AddPermissionsToRole(c.Request.Context(), grpcReq)
	if err != nil {
		logger.Error("failed to add permissions", slog.String("role_id", req.RoleId.String()), slog.String("error", err.Error()))
		h.handleGRPCError(c, err)
		return
	}

	status := models.OperationStatusFromProto(resp)

	logger.Info("permissions added successfully", slog.String("role_id", req.RoleId.String()), slog.Int("permissions_added", len(req.PermissionIDs)))

	c.JSON(http.StatusOK, gin.H{
		"status":  status.Success,
		"message": status.Message,
		"details": gin.H{
			"role_id":        req.RoleId.String(),
			"permission_ids": req.PermissionIDs,
			"timestamp":      status.Timestamp,
		},
	})
}
