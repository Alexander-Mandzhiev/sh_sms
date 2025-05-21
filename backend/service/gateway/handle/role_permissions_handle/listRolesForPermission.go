package role_permissions_handle

import (
	"backend/protos/gen/go/sso/role_permissions"
	"backend/service/gateway/models/sso"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) listRolesForPermission(c *gin.Context) {
	const op = "gateway.RolePermissions.ListRolesForPermission"
	logger := h.logger.With(slog.String("op", op))

	var req sso_models.ListRolesForPermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("failed to bind JSON", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if req.PermissionID.String() == "" || req.ClientID.String() == "" || req.AppID == 0 {
		logger.Error("missing required parameters", slog.String("permission_id", req.PermissionID.String()), slog.String("client_id", req.ClientID.String()), slog.Int("app_id", int(req.AppID)))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters 'permission_id', 'client_id' and 'app_id' are required"})
		return
	}

	grpcReq := &role_permissions.ListRolesRequest{
		PermissionId: req.PermissionID.String(),
		ClientId:     req.ClientID.String(),
		AppId:        req.AppID,
	}

	resp, err := h.service.ListRolesForPermission(c.Request.Context(), grpcReq)
	if err != nil {
		logger.Error("failed to list roles", slog.String("permission_id", req.PermissionID.String()), slog.String("client_id", req.ClientID.String()), slog.Int("app_id", int(req.AppID)), slog.String("error", err.Error()))
		h.handleGRPCError(c, err)
		return
	}

	rolesList := sso_models.RolesForPermissionListFromProto(resp)
	logger.Info("roles listed successfully", slog.String("permission_id", req.PermissionID.String()), slog.Int("roles_count", len(rolesList.RoleIDs)))
	c.JSON(http.StatusOK, gin.H{
		"data": rolesList.RoleIDs,
		"meta": gin.H{
			"permission_id": req.PermissionID,
			"count":         len(rolesList.RoleIDs),
		},
	})
}
