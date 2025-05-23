package auth_handle

import (
	"backend/protos/gen/go/auth"
	"backend/service/gateway/models/auth"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) listUserSessions(c *gin.Context) {
	const op = "auth_handler.listUserSessions"
	logger := h.logger.With(slog.String("op", op))

	var req auth_models.SessionFilter
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("failed to bind request", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	pbFilter := &auth.SessionFilter{
		UserId:     req.UserID.String(),
		ClientId:   req.ClientID.String(),
		AppId:      int32(req.AppID),
		ActiveOnly: req.ActiveOnly,
		Page:       int32(req.Page),
		Count:      int32(req.Count),
	}

	sessions, err := h.service.ListSessionsForUser(c.Request.Context(), pbFilter)
	if err != nil {
		logger.Error("failed to get sessions", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	response := auth_models.SessionsFromProto(sessions.Sessions)
	logger.Info("successfully retrieved sessions", slog.Int("count", len(sessions.Sessions)), slog.Int("page", req.Page))
	c.JSON(http.StatusOK, gin.H{"sessions": response, "count": len(response)})
}
