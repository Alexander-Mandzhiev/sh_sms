package auth_handle

import (
	"backend/protos/gen/go/auth"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
)

func (h *Handler) terminateSession(c *gin.Context) {
	const op = "grpc.handler.TerminateSession"
	logger := h.logger.With(slog.String("op", op))

	sessionID := c.Param("session_id")
	if sessionID == "" {
		logger.Error("session_id is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "session_id is required"})
		return
	}

	sessionUUID, err := uuid.Parse(sessionID)
	if err != nil {
		logger.Error("invalid session_id format", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session_id format"})
		return
	}

	pbRequest := &auth.SessionID{
		SessionId: sessionUUID.String(),
	}

	err = h.service.TerminateSession(c.Request.Context(), pbRequest)
	if err != nil {
		logger.Error("failed to terminate session", slog.Any("error", err), slog.String("session_id", sessionID))
		switch {
		case errors.Is(err, ErrSessionNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
		case errors.Is(err, ErrSessionAlreadyRevoked):
			c.JSON(http.StatusConflict, gin.H{"error": "session already revoked"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	logger.Info("session terminated successfully", slog.String("session_id", sessionID))
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "session terminated"})
}
