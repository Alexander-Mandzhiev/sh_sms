package attachments_handle

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) delete(c *gin.Context) {
	const op = "gateway.Attachments.Delete"
	logger := h.logger.With(slog.String("op", op))

	fileId := c.Param("file_id")
	if fileId == "" {
		logger.Error("File ID required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "File ID is required"})
		return
	}

	if err := h.service.DeleteAttachment(c.Request.Context(), fileId); err != nil {
		logger.Error("Deletion failed", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete attachment"})
		return
	}

	logger.Info("Attachment deleted", "file_id", fileId)
	c.Status(http.StatusNoContent)
}
