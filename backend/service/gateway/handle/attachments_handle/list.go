package attachments_handle

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) listByBook(c *gin.Context) {
	const op = "gateway.Attachments.ListByBook"
	logger := h.logger.With(slog.String("op", op))

	bookID, err := strconv.ParseInt(c.Param("book_id"), 10, 64)
	if err != nil {
		logger.Error("Invalid book ID", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	attachments, err := h.service.ListAttachmentsByBook(c.Request.Context(), bookID)
	if err != nil {
		logger.Error("Failed to list attachments", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list attachments"})
		return
	}

	logger.Info("Attachments listed successfully")
	c.JSON(http.StatusOK, attachments)
}
