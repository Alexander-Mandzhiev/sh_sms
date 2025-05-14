package roles_handle

import (
	"log/slog"
	"net/http"
	"strconv"

	"backend/protos/gen/go/sso/roles"
	"backend/service/gateway/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) updateRole(c *gin.Context) {
	const op = "gateway.Roles.Update"
	logger := h.logger.With(slog.String("op", op))

	id := c.Param("id")
	clientID := c.Query("client_id")
	appIDStr := c.Query("app_id")

	if id == "" || clientID == "" || appIDStr == "" {
		logger.Error("Missing required parameters", slog.String("role_id", id), slog.String("client_id", clientID), slog.String("app_id", appIDStr))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Parameters 'id', 'client_id' and 'app_id' are required",
		})
		return
	}

	appID, err := strconv.Atoi(appIDStr)
	if err != nil {
		logger.Error("Invalid app_id format", slog.String("app_id", appIDStr), slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid app_id format, must be integer"})
		return
	}

	var req models.UpdateRoleRequest
	if err = c.ShouldBindJSON(&req); err != nil {
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
		ClientId:    clientID,
		AppId:       int32(appID),
		Name:        req.Name,
		Description: req.Description,
		Level:       req.Level,
	}

	resp, err := h.service.UpdateRole(c.Request.Context(), grpcReq)
	if err != nil {
		logger.Error("Failed to update role", slog.String("role_id", id), slog.String("client_id", clientID), slog.Int("app_id", appID), slog.String("error", err.Error()))
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
