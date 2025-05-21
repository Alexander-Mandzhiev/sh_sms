package auth_handle

import (
	"backend/pkg/cookies"
	"backend/pkg/utils"
	"backend/protos/gen/go/auth"
	"backend/service/gateway/models/auth"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"time"
)

func (h *Handler) login(c *gin.Context) {
	const op = "auth_handler.Login"
	logger := h.logger.With(slog.String("op", op))

	var req auth_models.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("invalid request format", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := utils.ValidateUUID(req.ClientID); err != nil {
		logger.Error("invalid client id", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid client_id"})
		return
	}

	grpcReq := &auth.LoginRequest{
		ClientId: req.ClientID.String(),
		AppId:    int32(req.AppID),
		Login:    req.Login,
		Password: req.Password,
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	res, err := h.service.Login(ctx, grpcReq)
	if err != nil {
		logger.Error("login failed", "error", err)
		h.handleGRPCError(c, err)
		return
	}

	cookies.SetRefreshCookie(c.Writer, res.RefreshToken, cookies.DefaultConfig, h.cfg.TokenDuration)

	response, err := auth_models.AuthResponseFromProto(res)
	if err != nil {
		logger.Error("failed to convert auth response", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	logger.Info("successful login", "user_id", response.User.ID)
	c.JSON(http.StatusOK, response)
}
