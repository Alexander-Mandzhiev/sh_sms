package attachments_handle

import (
	library_models "backend/pkg/models/library"
	library "backend/protos/gen/go/library"
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

	includeDeleted := c.Query("include_deleted") == "true"

	res, err := h.service.ListAttachmentsByBook(c.Request.Context(), &library.ListAttachmentsByBookRequest{BookId: bookID, IncludeDeleted: includeDeleted})
	if err != nil {
		logger.Error("Failed to list attachments", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list attachments"})
		return
	}

	var attachments []*library_models.Attachment
	for _, attachment := range res.Attachments {
		att := library_models.AttachmentFromProto(attachment)
		attachments = append(attachments, att)
	}

	logger.Info("Attachments listed successfully")
	c.JSON(http.StatusOK, attachments)
}
