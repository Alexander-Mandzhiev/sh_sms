package role_permissions_handle

import (
	"log/slog"
	"net/http"

	"backend/protos/gen/go/sso/role_permissions"
	"backend/service/gateway/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) listPermissionsForRole(c *gin.Context) {
	const op = "gateway.RolePermissions.ListPermissionsForRole"
	logger := h.logger.With(slog.String("op", op))

	var req models.ListRolePermissionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("failed to bind JSON", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if req.RoleID == "" || req.ClientID == "" || req.AppID == 0 {
		logger.Warn("missing required fields", slog.String("role_id", req.RoleID), slog.String("client_id", req.ClientID), slog.Int("app_id", int(req.AppID)))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "role_id, client_id and app_id are required",
		})
		return
	}

	grpcReq := &role_permissions.ListPermissionsRequest{
		RoleId:   req.RoleID,
		ClientId: req.ClientID,
		AppId:    req.AppID,
	}

	resp, err := h.service.ListPermissionsForRole(c.Request.Context(), grpcReq)
	if err != nil {
		logger.Error("failed to list permissions", slog.String("role_id", req.RoleID), slog.String("client_id", req.ClientID), slog.Int("app_id", int(req.AppID)), slog.String("error", err.Error()))
		h.handleGRPCError(c, err)
		return
	}

	permissionsList := models.RolePermissionsListFromProto(resp)

	logger.Info("permissions listed successfully", slog.String("role_id", req.RoleID), slog.Int("permissions_count", len(permissionsList.PermissionIDs)))
	c.JSON(http.StatusOK, gin.H{
		"data": permissionsList.PermissionIDs,
		"meta": gin.H{
			"role_id": req.RoleID,
			"count":   len(permissionsList.PermissionIDs),
		},
	})
}
