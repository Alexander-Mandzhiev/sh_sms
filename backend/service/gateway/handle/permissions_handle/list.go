package permissions_handle

import (
	"backend/service/gateway/models/sso"
	"log/slog"
	"net/http"

	"backend/protos/gen/go/sso/permissions"
	"github.com/gin-gonic/gin"
)

func (h *Handler) list(c *gin.Context) {
	const op = "gateway.Permissions.List"
	logger := h.logger.With(slog.String("op", op))

	var req sso_models.ListPermissionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Invalid request format", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	grpcReq := &permissions.ListRequest{
		AppId: req.AppID,
		Page:  int32(req.Pagination.Page),
		Count: int32(req.Pagination.Count),
	}

	if req.Filters.CodeFilter != nil {
		grpcReq.CodeFilter = req.Filters.CodeFilter
	}
	if req.Filters.Category != nil {
		grpcReq.Category = req.Filters.Category
	}
	if req.Filters.ActiveOnly != nil {
		grpcReq.ActiveOnly = req.Filters.ActiveOnly
	}

	resp, err := h.service.ListPermissions(c.Request.Context(), grpcReq)
	if err != nil {
		logger.Error("Failed to list permissions", slog.Int("app_id", int(req.AppID)), slog.String("error", err.Error()))
		h.handleGRPCError(c, err)
		return
	}

	permissionsList := make([]*sso_models.Permission, 0, len(resp.Permissions))
	for _, protoPerm := range resp.Permissions {
		perm, err := sso_models.PermissionFromProto(protoPerm)
		if err != nil {
			logger.Error("Proto conversion failed", slog.String("error", err.Error()))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		permissionsList = append(permissionsList, perm)
	}

	logger.Info("Permissions listed successfully", slog.Int("count", len(permissionsList)), slog.Int("total", int(resp.TotalCount)))
	c.JSON(http.StatusOK, gin.H{
		"data": permissionsList,
		"meta": gin.H{
			"total_count":    resp.TotalCount,
			"current_page":   resp.CurrentPage,
			"count_per_page": req.Pagination.Count,
		},
	})
}
