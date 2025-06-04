package attachments_handle

import (
	"backend/pkg/models/library"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

// update обновляет вложение
func (h *Handler) update(c *gin.Context) {
	const op = "gateway.Attachments.Update"
	logger := h.logger.With(slog.String("op", op))

	var updateReq library_models.UpdateAttachmentRequest
	if err := c.ShouldBindJSON(&updateReq); err != nil {
		logger.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	req := updateReq.UpdateAttachmentRequestToProto()
	res, err := h.service.UpdateAttachment(c.Request.Context(), req)
	if err != nil {
		logger.Error("Failed to update attachment", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update attachment"})
		return
	}

	logger.Info("Attachment updated successfully")
	c.JSON(http.StatusOK, res)
}
