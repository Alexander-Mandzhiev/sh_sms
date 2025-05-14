package roles_handle

import (
	"log/slog"
	"net/http"
	"strconv"

	"backend/protos/gen/go/sso/roles"
	"github.com/gin-gonic/gin"
)

func (h *Handler) deleteRole(c *gin.Context) {
	const op = "gateway.Roles.Delete"
	logger := h.logger.With(slog.String("op", op))

	id := c.Param("id")
	clientID := c.Query("client_id")
	appIDStr := c.Query("app_id")
	permanentStr := c.DefaultQuery("permanent", "false")

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

	permanent, err := strconv.ParseBool(permanentStr)
	if err != nil {
		logger.Error("Invalid permanent flag format", slog.String("permanent", permanentStr), slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid permanent value, must be boolean"})
		return
	}

	grpcReq := &roles.DeleteRequest{
		Id:        id,
		ClientId:  clientID,
		AppId:     int32(appID),
		Permanent: &permanent,
	}

	resp, err := h.service.DeleteRole(c.Request.Context(), grpcReq)
	if err != nil {
		logger.Error("Failed to delete role", slog.String("role_id", id), slog.String("client_id", clientID), slog.Int("app_id", appID), slog.Bool("permanent", permanent), slog.String("error", err.Error()))
		h.handleGRPCError(c, err)
		return
	}

	logger.Info("Role deleted successfully", slog.String("role_id", id), slog.Bool("permanent", permanent))
	c.JSON(http.StatusOK, gin.H{"success": resp.Success})
}
