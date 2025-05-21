package auth_handle

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func (h *Handler) handleGRPCError(c *gin.Context, err error) {
	st, ok := status.FromError(err)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	switch st.Code() {
	case codes.InvalidArgument:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
	case codes.NotFound:
		c.JSON(http.StatusNotFound, gin.H{"error": "resource not found"})
	case codes.Unauthenticated:
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication failed"})
	case codes.PermissionDenied:
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
	case codes.AlreadyExists:
		c.JSON(http.StatusConflict, gin.H{"error": "resource conflict"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}
