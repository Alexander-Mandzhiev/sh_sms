package attachments_handle

import (
	library_models "backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) get(c *gin.Context) {
	const op = "gateway.Attachments.Get"
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

	res, err := h.service.GetAttachment(c.Request.Context(), &library.GetAttachmentRequest{
		BookId: bookID,
		Format: format,
	})
	if err != nil {
		logger.Error("Failed to get attachment", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get attachment"})
		return
	}

	response := library_models.AttachmentFromProto(res)

	logger.Info("Attachment retrieved successfully")
	c.JSON(http.StatusOK, response)
}
