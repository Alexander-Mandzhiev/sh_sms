package file_formats_handle

import (
	library "backend/protos/gen/go/library"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) exist(c *gin.Context) {
	const op = "gateway.FileFormat.Exist"
	logger := h.logger.With(slog.String("op", op))
	format := c.Param("format")
	if format == "" {
		logger.Error("Missing required id", slog.String("id", format))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters id are required"})
		return
	}

	res, err := h.service.FileFormatExists(c.Request.Context(), &library.FileFormatExistsRequest{Format: format})
	if err != nil {
		logger.Error("Error existing file formats", slog.String("format", format), slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Info("File format existing")
	c.JSON(http.StatusOK, res)
}
