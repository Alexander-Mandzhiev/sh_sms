package user_roles_handle

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"net/http"
)

func (h *Handler) handleGRPCError(c *gin.Context, err error) {
	logger := h.logger.With(slog.String("op", "handleGRPCError"))
	st, ok := status.FromError(err)
	if !ok {
		logger.Error("Unknown error type", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	switch st.Code() {
	case codes.AlreadyExists:
		logger.Warn("Conflict", slog.String("message", st.Message()))
		c.JSON(http.StatusConflict, gin.H{"error": "Role already assigned to user"})
	case codes.NotFound:
		logger.Warn("Not found", slog.String("message", st.Message()))
		c.JSON(http.StatusNotFound, gin.H{"error": "User or role not found"})
	case codes.InvalidArgument:
		logger.Warn("Invalid request", slog.String("message", st.Message()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameters", "details": st.Message()})
	case codes.PermissionDenied:
		logger.Warn("Permission denied", slog.String("message", st.Message()))
		c.JSON(http.StatusForbidden, gin.H{"error": "Operation not permitted"})
	case codes.FailedPrecondition:
		logger.Warn("Precondition failed", slog.String("message", st.Message()))
		c.JSON(http.StatusPreconditionFailed, gin.H{"error": "Precondition check failed", "details": st.Message()})
	case codes.Unavailable:
		logger.Error("Service unavailable", slog.String("message", st.Message()))
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Service temporarily unavailable"})
	default:
		logger.Error("Unexpected error", slog.String("code", st.Code().String()), slog.String("message", st.Message()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
}
