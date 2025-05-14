package roles_handle

import (
	"log/slog"
	"net/http"
	"strconv"

	"backend/protos/gen/go/sso/roles"
	"backend/service/gateway/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) restoreRole(c *gin.Context) {
	const op = "gateway.Roles.Restore"
	logger := h.logger.With(slog.String("op", op))

	id := c.Param("id")
	clientID := c.Query("client_id")
	appIDStr := c.Query("app_id")

	if id == "" || clientID == "" || appIDStr == "" {
		logger.Error("Missing required parameters", slog.String("role_id", id), slog.String("client_id", clientID), slog.String("app_id", appIDStr))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters 'id', 'client_id' and 'app_id' are required"})
		return
	}

	appID, err := strconv.Atoi(appIDStr)
	if err != nil {
		logger.Error("Invalid app_id format", slog.String("app_id", appIDStr), slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid app_id format, must be integer"})
		return
	}

	grpcReq := &roles.RestoreRequest{
		Id:       id,
		ClientId: clientID,
		AppId:    int32(appID),
	}

	resp, err := h.service.RestoreRole(c.Request.Context(), grpcReq)
	if err != nil {
		logger.Error("Failed to restore role", slog.String("role_id", id), slog.String("client_id", clientID), slog.Int("app_id", appID), slog.String("error", err.Error()))
		h.handleGRPCError(c, err)
		return
	}

	role, err := models.RoleFromProto(resp)
	if err != nil {
		logger.Error("Proto conversion failed", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	logger.Info("Role restored successfully", slog.String("role_id", role.ID), slog.String("client_id", role.ClientID), slog.Int("app_id", int(role.AppID)))
	c.JSON(http.StatusOK, role)
}
