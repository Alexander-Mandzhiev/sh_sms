package attachments_handle

import (
	library "backend/protos/gen/go/library"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) deleteFile(c *gin.Context) {
	const op = "gateway.Attachments.DeleteFile"
	logger := h.logger.With(slog.String("op", op))

	fileURL := c.Param("file_url")
	if fileURL == "" {
		logger.Error("File URL parameter is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "File URL parameter is required"})
		return
	}

	_, err := h.service.DeleteFile(c.Request.Context(), &library.DeleteFileRequest{
		FileUrl: fileURL,
	})
	if err != nil {
		logger.Error("Failed to delete file", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file"})
		return
	}

	logger.Info("File deleted successfully")
	c.Status(http.StatusNoContent)
}
