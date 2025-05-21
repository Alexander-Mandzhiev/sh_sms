package users_handle

import (
	"backend/protos/gen/go/sso/users"
	"backend/service/gateway/models/sso"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) createUser(c *gin.Context) {
	const op = "gateway.Users.Create"
	logger := h.logger.With(slog.String("op", op))

	var req sso_models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	grpcReq := &users.CreateRequest{
		ClientId: req.ClientID.String(),
		Email:    req.Email,
		Password: req.Password,
		FullName: req.FullName,
		Phone:    req.Phone,
	}

	resp, err := h.service.CreateUser(c.Request.Context(), grpcReq)
	if err != nil {
		logger.Error("Failed to create user", "error", err)
		h.handleGRPCError(c, err)
		return
	}

	user, err := sso_models.UserFromProto(resp)
	if err != nil {
		logger.Error("Failed to convert proto user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	logger.Info("User created successfully", "user_id", user.ID)
	c.JSON(http.StatusCreated, user)
}
