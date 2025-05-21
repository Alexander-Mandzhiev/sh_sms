package user_roles_handle

import (
	"backend/pkg/utils"
	user_roles "backend/protos/gen/go/sso/users_roles"
	"backend/service/gateway/models/sso"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) listForRole(c *gin.Context) {
	const op = "gateway.UserRoles.ListForRole"
	logger := h.logger.With(slog.String("op", op))

	var req sso_models.RoleUsersListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to parse request", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if req.RoleID == "" || req.ClientID == "" || req.AppID == 0 {
		logger.Warn("Missing required fields", slog.String("role_id", req.RoleID), slog.String("client_id", req.ClientID), slog.Int("app_id", int(req.AppID)))
		c.JSON(http.StatusBadRequest, gin.H{"error": "role_id, client_id and app_id are required"})
		return
	}

	if err := utils.ValidatePagination(int(req.Page), int(req.Count)); err != nil {
		logger.Warn("Invalid pagination", slog.Int("page", int(req.Page)), slog.Int("count", int(req.Count)))
		c.JSON(http.StatusBadRequest, gin.H{"error": "page must be >=1, count between 1-1000"})
		return
	}

	grpcReq := &user_roles.RoleRequest{
		RoleId:   req.RoleID,
		ClientId: req.ClientID,
		AppId:    req.AppID,
		Page:     req.Page,
		Count:    req.Count,
	}

	resp, err := h.service.ListForRole(c.Request.Context(), grpcReq)
	if err != nil {
		logger.Error("Failed to list users for role", slog.String("role_id", req.RoleID), slog.String("client_id", req.ClientID), slog.String("error", err.Error()))
		h.handleGRPCError(c, err)
		return
	}

	response := sso_models.UserRolesListFromProto(resp)
	logger.Info("Users listed for role", slog.String("role_id", req.RoleID), slog.Int("users_count", len(response.Assignments)))
	c.JSON(http.StatusOK, gin.H{
		"data": response.Assignments,
		"meta": gin.H{
			"role_id":        req.RoleID,
			"total_count":    response.TotalCount,
			"current_page":   response.CurrentPage,
			"count_per_page": req.Count,
			"app_id":         response.AppID,
		},
	})
}
