package roles_handle

import (
	"backend/service/gateway/models/sso"
	"log/slog"
	"net/http"

	"backend/protos/gen/go/sso/roles"
	"github.com/gin-gonic/gin"
)

func (h *Handler) listRoles(c *gin.Context) {
	const op = "gateway.Roles.List"
	logger := h.logger.With(slog.String("op", op))

	var req sso_models.ListRolesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Invalid request format", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	grpcReq := &roles.ListRequest{
		ClientId: req.ClientID,
		AppId:    int32(req.AppID),
		Page:     int32(req.Pagination.Page),
		Count:    int32(req.Pagination.Count),
	}

	if req.Filters.Name != nil {
		grpcReq.NameFilter = req.Filters.Name
	}
	if req.Filters.Level != nil {
		grpcReq.LevelFilter = req.Filters.Level
	}
	if req.Filters.ActiveOnly != nil {
		grpcReq.ActiveOnly = req.Filters.ActiveOnly
	}

	resp, err := h.service.ListRoles(c.Request.Context(), grpcReq)
	if err != nil {
		logger.Error("Failed to list roles", slog.String("client_id", req.ClientID), slog.Int("app_id", req.AppID), slog.String("error", err.Error()))
		h.handleGRPCError(c, err)
		return
	}

	rolesList := make([]*sso_models.Role, 0, len(resp.Roles))
	for _, protoRole := range resp.Roles {
		role, err := sso_models.RoleFromProto(protoRole)
		if err != nil {
			logger.Error("Proto conversion failed", slog.String("error", err.Error()))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		rolesList = append(rolesList, role)
	}

	logger.Info("Roles listed successfully", slog.Int("count", len(rolesList)), slog.Int("total", int(resp.TotalCount)))

	c.JSON(http.StatusOK, gin.H{
		"data": rolesList,
		"meta": gin.H{
			"total_count":    resp.TotalCount,
			"current_page":   resp.CurrentPage,
			"count_per_page": req.Pagination.Count,
		},
	})
}
