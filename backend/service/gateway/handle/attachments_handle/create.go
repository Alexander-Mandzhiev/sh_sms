package attachments_handle

import (
	"backend/pkg/models/library"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) create(c *gin.Context) {
	const op = "gateway.Attachments.Create"
	logger := h.logger.With(slog.String("op", op))

	var reqCreate library_models.CreateAttachmentRequest
	if err := c.ShouldBindJSON(&reqCreate); err != nil {
		logger.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	req, err := reqCreate.CreateAttachmentRequestToProto()
	if err != nil {
		logger.Error("Failed to create attachment request", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	res, err := h.service.CreateAttachment(c.Request.Context(), req)
	if err != nil {
		logger.Error("Failed to create attachment", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create attachment"})
		return
	}
	response := library_models.AttachmentFromProto(res)
	logger.Info("Attachment created successfully")
	c.JSON(http.StatusCreated, response)
}
