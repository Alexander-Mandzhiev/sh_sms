package role_permissions_handle

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func (h *Handler) handleGRPCError(c *gin.Context, err error) {
	st, ok := status.FromError(err)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	switch st.Code() {
	case codes.NotFound:
		c.JSON(http.StatusNotFound, gin.H{"error": "role or permission not found"})
	case codes.AlreadyExists:
		c.JSON(http.StatusConflict, gin.H{"error": "permission already assigned to role"})
	case codes.InvalidArgument:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request parameters", "details": st.Message()})
	case codes.PermissionDenied:
		c.JSON(http.StatusForbidden, gin.H{"error": "operation not permitted"})
	case codes.FailedPrecondition:
		c.JSON(http.StatusPreconditionFailed, gin.H{"error": "precondition check failed", "details": st.Message()})
	case codes.Unavailable:
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "service temporarily unavailable"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}
