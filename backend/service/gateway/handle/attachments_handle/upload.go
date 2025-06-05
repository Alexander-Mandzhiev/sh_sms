package attachments_handle

import (
	"backend/pkg/models/library"
	"github.com/gin-gonic/gin"
	"log/slog"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

func (h *Handler) uploadFile(c *gin.Context) {
	const op = "gateway.Attachments.Upload"
	logger := h.logger.With(slog.String("op", op))

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		logger.Error("File required", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}
	defer file.Close()

	bookID, err := strconv.ParseInt(c.PostForm("book_id"), 10, 64)
	if err != nil {
		logger.Error("Invalid book ID", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	format := determineFileFormat(header.Filename, header.Header.Get("Content-Type"))
	attachment, err := h.service.CreateAttachment(c.Request.Context(), library_models.FileMetadata{BookID: bookID, Format: format}, file)

	if err != nil {
		logger.Error("Upload failed", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "File upload failed"})
		return
	}

	logger.Info("File uploaded successfully", "book_id", bookID, "file_id", attachment.FileID)
	c.JSON(http.StatusCreated, attachment)
}

func determineFileFormat(filename, contentType string) string {
	if contentType != "" {
		if exts, _ := mime.ExtensionsByType(contentType); len(exts) > 0 {
			return strings.TrimPrefix(exts[0], ".")
		}
	}
	if ext := filepath.Ext(filename); ext != "" {
		return strings.TrimPrefix(ext, ".")
	}
	return "unknown"
}
