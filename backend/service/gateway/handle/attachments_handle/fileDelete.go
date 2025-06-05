package attachments_handle

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

func (h *Handler) deleteFile(c *gin.Context) {
	const op = "gateway.Attachments.DownloadFile"
	logger := h.logger.With(slog.String("op", op))
	_ = logger
}
