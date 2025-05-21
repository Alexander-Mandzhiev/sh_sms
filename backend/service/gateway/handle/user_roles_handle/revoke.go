package user_roles_handle

import (
	"backend/service/gateway/models/sso"
	"log/slog"
	"net/http"
	"time"

	"backend/protos/gen/go/sso/users_roles"
	"github.com/gin-gonic/gin"
)

func (h *Handler) revokeRole(c *gin.Context) {
	const op = "gateway.UserRoles.RevokeRole"
	logger := h.logger.With(slog.String("op", op))

	var req sso_models.RevokeRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to parse request", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if req.UserID == "" || req.RoleID == "" || req.ClientID == "" || req.AppID == 0 {
		logger.Warn("Missing required fields", slog.String("user_id", req.UserID), slog.String("role_id", req.RoleID), slog.String("client_id", req.ClientID), slog.Int("app_id", int(req.AppID)))
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id, role_id, client_id and app_id are required"})
		return
	}

	grpcReq := &user_roles.RevokeRequest{
		UserId:   req.UserID,
		RoleId:   req.RoleID,
		ClientId: req.ClientID,
		AppId:    req.AppID,
	}

	resp, err := h.service.RevokeRole(c.Request.Context(), grpcReq)
	if err != nil {
		logger.Error("Failed to revoke role", slog.String("user_id", req.UserID), slog.String("role_id", req.RoleID), slog.String("error", err.Error()))
		h.handleGRPCError(c, err)
		return
	}

	revokeResponse := sso_models.RevokeResponseFromProto(resp)
	logger.Info("Role revoked successfully", slog.String("user_id", req.UserID), slog.String("role_id", req.RoleID))
	c.JSON(http.StatusOK, gin.H{
		"success":    revokeResponse.Success,
		"revoked_at": revokeResponse.RevokedAt.Format(time.RFC3339),
	})
}
