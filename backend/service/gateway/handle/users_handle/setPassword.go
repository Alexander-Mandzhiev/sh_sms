package users_handle

import (
	"backend/service/gateway/models/sso"
	"log/slog"
	"net/http"

	"backend/protos/gen/go/sso/users"
	"github.com/gin-gonic/gin"
)

func (h *Handler) setUserPassword(c *gin.Context) {
	const op = "gateway.Users.SetPassword"
	logger := h.logger.With(slog.String("op", op))

	id := c.Param("id")
	clientID := c.Query("client_id")

	if id == "" || clientID == "" {
		logger.Error("Missing required parameters")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Both 'id' path parameter and 'client_id' query parameter are required",
		})
		return
	}

	var reqBody sso_models.SetPasswordRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		logger.Error("Invalid request body", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	grpcReq := &users.SetPasswordRequest{
		Id:          id,
		ClientId:    clientID,
		NewPassword: reqBody.NewPassword,
	}

	resp, err := h.service.SetPassword(c.Request.Context(), grpcReq)
	if err != nil {
		logger.Error("Failed to set password", slog.String("user_id", id), slog.String("client_id", clientID), slog.String("error", err.Error()))
		h.handleGRPCError(c, err)
		return
	}

	logger.Info("Password updated successfully", slog.String("user_id", id), slog.String("client_id", clientID))
	c.JSON(http.StatusOK, gin.H{"success": resp.Success})
}
