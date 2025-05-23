package user_roles_handle

import (
	user_roles "backend/protos/gen/go/sso/users_roles"
	"backend/service/gateway/models/sso"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"net/http"
)

func (h *Handler) assignRole(c *gin.Context) {
	const op = "gateway.UserRoles.AssignRole"
	logger := h.logger.With(slog.String("op", op))

	var req sso_models.AssignRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to bind JSON", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if req.UserID == "" || req.RoleID == "" || req.ClientID == "" || req.AppID == 0 || req.AssignedBy == "" {
		logger.Warn("Missing required fields", slog.String("user_id", req.UserID), slog.String("role_id", req.RoleID), slog.String("client_id", req.ClientID), slog.Int("app_id", int(req.AppID)), slog.String("assigned_by", req.AssignedBy))
		c.JSON(http.StatusBadRequest, gin.H{"error": "All required fields must be provided"})
		return
	}

	var expiresAtPB *timestamppb.Timestamp
	if req.ExpiresAt != nil {
		expiresAtPB = timestamppb.New(*req.ExpiresAt)
	}

	grpcReq := &user_roles.AssignRequest{
		UserId:     req.UserID,
		RoleId:     req.RoleID,
		ClientId:   req.ClientID,
		AppId:      req.AppID,
		AssignedBy: req.AssignedBy,
		ExpiresAt:  expiresAtPB,
	}

	resp, err := h.service.AssignRole(c.Request.Context(), grpcReq)
	if err != nil {
		logger.Error("Failed to assign role", slog.String("user_id", req.UserID), slog.String("role_id", req.RoleID), slog.String("error", err.Error()))
		h.handleGRPCError(c, err)
		return
	}

	userRole := sso_models.UserRoleFromProto(resp)
	logger.Info("Role assigned successfully", slog.String("user_id", userRole.UserID), slog.String("role_id", userRole.RoleID))
	c.JSON(http.StatusCreated, userRole)
}
