package permissions_handle

import (
	"log/slog"
	"net/http"
	"strconv"

	"backend/protos/gen/go/sso/permissions"
	"github.com/gin-gonic/gin"
)

func (h *Handler) delete(c *gin.Context) {
	const op = "gateway.Permissions.Delete"
	logger := h.logger.With(slog.String("op", op))

	id := c.Param("id")
	appIDStr := c.Query("app_id")
	permanentStr := c.DefaultQuery("permanent", "false")

	if id == "" || appIDStr == "" {
		logger.Error("Missing required parameters", slog.String("permission_id", id), slog.String("app_id", appIDStr))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters 'id' and 'app_id' are required"})
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

	grpcReq := &permissions.DeleteRequest{
		Id:        id,
		AppId:     int32(appID),
		Permanent: &permanent,
	}

	resp, err := h.service.DeletePermission(c.Request.Context(), grpcReq)
	if err != nil {
		logger.Error("Failed to delete permission", slog.String("permission_id", id), slog.Int("app_id", appID), slog.Bool("permanent", permanent), slog.String("error", err.Error()))
		h.handleGRPCError(c, err)
		return
	}

	logger.Info("Permission deleted successfully", slog.String("permission_id", id), slog.Bool("permanent", permanent))

	c.JSON(http.StatusOK, gin.H{"success": resp.Success})
}
