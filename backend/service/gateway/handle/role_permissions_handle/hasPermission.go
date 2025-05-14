package role_permissions_handle

import (
	"log/slog"
	"net/http"
	"time"

	"backend/protos/gen/go/sso/role_permissions"
	"backend/service/gateway/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) hasPermission(c *gin.Context) {
	const op = "gateway.RolePermissions.HasPermission"
	logger := h.logger.With(slog.String("op", op))

	var req models.HasPermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("failed to bind JSON", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if req.RoleID.String() == "" || req.PermissionID.String() == "" || req.ClientID.String() == "" || req.AppID == 0 {
		logger.Error("missing required parameters", slog.String("role_id", req.RoleID.String()), slog.String("permission_id", req.PermissionID.String()), slog.String("client_id", req.ClientID.String()), slog.Int("app_id", int(req.AppID)))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters 'role_id', 'permission_id', 'client_id' and 'app_id' are required"})
		return
	}

	grpcReq := &role_permissions.HasPermissionRequest{
		RoleId:       req.RoleID.String(),
		PermissionId: req.PermissionID.String(),
		ClientId:     req.ClientID.String(),
		AppId:        req.AppID,
	}

	resp, err := h.service.HasPermission(c.Request.Context(), grpcReq)
	if err != nil {
		logger.Error("failed to check permission", slog.String("role_id", req.RoleID.String()), slog.String("permission_id", req.PermissionID.String()), slog.String("error", err.Error()))
		h.handleGRPCError(c, err)
		return
	}

	result := models.HasPermissionResponseFromProto(resp)

	logger.Info("permission check completed", slog.Bool("result", result.HasPermission), slog.String("permission_id", req.PermissionID.String()))
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"has_permission": result.HasPermission,
			"checked_at":     result.CheckedAt.Format(time.RFC3339),
		},
		"meta": gin.H{
			"role_id":       req.RoleID.String(),
			"permission_id": req.PermissionID.String(),
			"client_id":     req.ClientID.String(),
			"app_id":        req.AppID,
		},
	})
}
