// attachments_handle.go
package attachments_handle

import (
	"github.com/gin-gonic/gin"
	"io"
	"log/slog"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"
)

func (h *Handler) get(c *gin.Context) {
	const op = "gateway.Attachments.Get"
	logger := h.logger.With(slog.String("op", op))

	bookID, err := strconv.ParseInt(c.Param("book_id"), 10, 64)
	if err != nil {
		logger.Error("Invalid book ID", "error", err, "input", c.Param("book_id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	format := c.Query("format")
	if format == "" {
		logger.Error("Format parameter is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format parameter is required"})
		return
	}

	logger = logger.With(slog.Int64("book_id", bookID), slog.String("format", format))
	logger.Debug("Processing attachment request")

	file, filePath, size, err := h.service.GetAttachment(c.Request.Context(), bookID, format)
	if err != nil {
		logger.Error("Failed to get attachment", "error", err, "details", "service call failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get attachment"})
		return
	}
	defer file.Close()

	logger.Debug("Attachment retrieved", "file_path", filePath, "file_size", size)

	contentType := "application/octet-stream"
	if ext := filepath.Ext(filePath); ext != "" {
		if ct := mime.TypeByExtension(ext); ct != "" {
			contentType = ct
		}
	}

	filename := filepath.Base(filePath)

	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", contentType)
	c.Header("Content-Length", strconv.FormatInt(size, 10))

	if _, err = io.Copy(c.Writer, file); err != nil {
		logger.Error("Failed to send file", "error", err, "details", "io.Copy operation failed")
		return
	}

	logger.Info("Attachment file sent successfully", "book_id", bookID, "format", format, "file_name", filename)
}
