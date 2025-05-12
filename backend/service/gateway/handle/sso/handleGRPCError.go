package sso_handle

import (
	"connectrpc.com/connect"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) handleGRPCError(c *gin.Context, err error, logMsg string) {
	grpcCode := connect.CodeOf(err)

	h.logger.Error(logMsg,
		slog.String("error", err.Error()),
		slog.String("code", grpcCode.String()),
	)

	switch grpcCode {
	case connect.CodeAlreadyExists:
		c.JSON(http.StatusConflict, gin.H{
			"error": "user with this email already exists",
		})
	case connect.CodeInvalidArgument:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input data",
		})
	case connect.CodePermissionDenied:
		c.JSON(http.StatusForbidden, gin.H{
			"error": "operation not permitted",
		})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
	}
}
