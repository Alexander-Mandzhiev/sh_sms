package users_handle

import (
	"backend/protos/gen/go/sso/users"
	"backend/service/gateway/models/sso"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) updateUser(c *gin.Context) {
	const op = "gateway.Users.Update"
	logger := h.logger.With(slog.String("op", op))

	id := c.Param("id")
	if id == "" {
		logger.Error("Missing user ID in URL")
		c.JSON(http.StatusBadRequest, gin.H{"error": "user ID is required"})
		return
	}

	var req sso_models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if req.ClientID == "" {
		logger.Error("Missing client ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "client_id is required"})
		return
	}

	grpcReq := &users.UpdateRequest{
		Id:       id,
		ClientId: req.ClientID,
		Email:    req.Email,
		FullName: req.FullName,
		Phone:    req.Phone,
	}

	resp, err := h.service.UpdateUser(c.Request.Context(), grpcReq)
	if err != nil {
		logger.Error("Failed to update user", "user_id", id, "client_id", req.ClientID, "error", err)
		h.handleGRPCError(c, err)
		return
	}

	user, err := sso_models.UserFromProto(resp)
	if err != nil {
		logger.Error("Proto conversion failed", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	logger.Info("User updated successfully", "user_id", user.ID)
	c.JSON(http.StatusOK, user)
}
