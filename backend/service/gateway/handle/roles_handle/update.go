package roles_handle

import (
	"backend/protos/gen/go/sso/roles"
	"backend/service/gateway/models"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) updateRole(c *gin.Context) {
	const op = "gateway.Roles.Update"
	logger := h.logger.With(slog.String("op", op))

	id := c.Param("id")
	if id == "" {
		logger.Error("Missing required parameters", slog.String("role_id", id))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters 'id', 'client_id' and 'app_id' are required"})
		return
	}

	var req models.UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to bind JSON", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if req.Name == nil && req.Description == nil && req.Level == nil {
		logger.Warn("No fields provided for update")
		c.JSON(http.StatusBadRequest, gin.H{"error": "At least one field (name, description or level) must be provided"})
		return
	}

	grpcReq := &roles.UpdateRequest{
		Id:          id,
		ClientId:    req.ClientID,
		AppId:       req.AppID,
		Name:        req.Name,
		Description: req.Description,
		Level:       req.Level,
	}

	resp, err := h.service.UpdateRole(c.Request.Context(), grpcReq)
	if err != nil {
		logger.Error("Failed to update role", slog.String("role_id", id), slog.String("client_id", req.ClientID), slog.Int("app_id", int(req.AppID)), slog.String("error", err.Error()))
		h.handleGRPCError(c, err)
		return
	}

	role, err := models.RoleFromProto(resp)
	if err != nil {
		logger.Error("Proto conversion failed", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	logger.Info("Role updated successfully", slog.String("role_id", role.ID), slog.String("client_id", role.ClientID), slog.Int("app_id", int(role.AppID)))
	c.JSON(http.StatusOK, role)
}
