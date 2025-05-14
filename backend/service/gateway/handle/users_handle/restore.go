package users_handle

import (
	"backend/protos/gen/go/sso/users"
	"backend/service/gateway/models"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) restoreUser(c *gin.Context) {
	const op = "gateway.Users.Restore"
	logger := h.logger.With(slog.String("op", op))

	id := c.Param("id")
	clientID := c.Query("client_id")

	if id == "" || clientID == "" {
		logger.Error("Missing required parameters")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Both user ID and client ID are required"})
		return
	}

	grpcReq := &users.RestoreRequest{
		Id:       id,
		ClientId: clientID,
	}

	resp, err := h.service.RestoreUser(c.Request.Context(), grpcReq)
	if err != nil {
		logger.Error("Failed to restore user", slog.String("user_id", id), slog.String("client_id", clientID), slog.String("error", err.Error()))
		h.handleGRPCError(c, err)
		return
	}

	user, err := models.UserFromProto(resp)
	if err != nil {
		logger.Error("Proto conversion failed", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	logger.Info("User restored successfully", slog.String("user_id", user.ID.String()))
	c.JSON(http.StatusOK, user)
}
