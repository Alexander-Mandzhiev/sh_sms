package attachments_handle

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) delete(c *gin.Context) {
	const op = "gateway.Attachments.Delete"
	logger := h.logger.With(slog.String("op", op))

	bookID, err := strconv.ParseInt(c.Param("book_id"), 10, 64)
	if err != nil {
		logger.Error("Invalid book ID", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	format := c.Query("format")
	if format == "" {
		logger.Error("Format parameter is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format parameter is required"})
		return
	}

	_, err = h.service.DeleteAttachment(c.Request.Context(), bookID, format)
	if err != nil {
		logger.Error("Failed to delete attachment", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete attachment"})
		return
	}

	logger.Info("Attachment deleted successfully")
	c.Status(http.StatusNoContent)
}
