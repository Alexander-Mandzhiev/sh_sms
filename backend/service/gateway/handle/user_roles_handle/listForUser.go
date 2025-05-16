package user_roles_handle

import (
	"backend/pkg/utils"
	user_roles "backend/protos/gen/go/sso/users_roles"
	"log/slog"
	"net/http"

	"backend/service/gateway/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) listForUser(c *gin.Context) {
	const op = "gateway.UserRoles.ListForUser"
	logger := h.logger.With(slog.String("op", op))

	var req models.UserRolesListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to parse request", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := utils.ValidatePagination(int(req.Page), int(req.Count)); err != nil {
		logger.Warn("Invalid pagination", slog.Int("page", int(req.Page)), slog.Int("count", int(req.Count)))
		c.JSON(http.StatusBadRequest, gin.H{"error": "page must be >=1, count between 1-1000"})
		return
	}

	grpcReq := &user_roles.UserRequest{
		UserId:   req.UserID,
		ClientId: req.ClientID,
		AppId:    req.AppID,
		Page:     req.Page,
		Count:    req.Count,
	}

	resp, err := h.service.ListForUser(c.Request.Context(), grpcReq)
	if err != nil {
		logger.Error("Failed to list user roles", slog.String("user_id", req.UserID), slog.String("client_id", req.ClientID), slog.Int("app_id", int(req.AppID)), slog.String("error", err.Error()))
		h.handleGRPCError(c, err)
		return
	}

	response := models.UserRolesListFromProto(resp)
	logger.Info("Roles listed successfully", slog.String("user_id", req.UserID), slog.Int("roles_count", len(response.Assignments)))
	c.JSON(http.StatusOK, gin.H{
		"data": response.Assignments,
		"meta": gin.H{
			"total_count":    response.TotalCount,
			"current_page":   response.CurrentPage,
			"count_per_page": req.Count,
			"app_id":         response.AppID,
		},
	})
}
