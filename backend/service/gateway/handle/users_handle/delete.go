package users_handle

import (
	"log/slog"
	"net/http"
	"strconv"

	"backend/protos/gen/go/sso/users"
	"github.com/gin-gonic/gin"
)

func (h *Handler) deleteUser(c *gin.Context) {
	const op = "gateway.Users.Delete"
	logger := h.logger.With(slog.String("op", op))

	id := c.Param("id")
	clientID := c.Query("client_id")
	permanentStr := c.DefaultQuery("permanent", "false")

	if id == "" || clientID == "" {
		logger.Error("Missing required parameters")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Both 'id' and 'client_id' parameters are required"})
		return
	}

	permanent, err := strconv.ParseBool(permanentStr)
	if err != nil {
		logger.Error("Invalid permanent parameter", slog.String("value", permanentStr), slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid permanent parameter, must be boolean"})
		return
	}

	grpcReq := &users.DeleteRequest{
		Id:        id,
		ClientId:  clientID,
		Permanent: permanent,
	}

	resp, err := h.service.DeleteUser(c.Request.Context(), grpcReq)
	if err != nil {
		logger.Error("Failed to delete user", slog.String("user_id", id), slog.String("client_id", clientID), slog.String("error", err.Error()))
		h.handleGRPCError(c, err)
		return
	}

	logger.Info("User deleted successfully", slog.String("user_id", id), slog.Bool("permanent", permanent))
	c.JSON(http.StatusOK, gin.H{"success": resp.Success})
}
