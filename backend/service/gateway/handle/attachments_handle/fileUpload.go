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
	const op = "gateway.Attachments.UploadFile"
	logger := h.logger.With(slog.String("op", op))

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		logger.Error("Failed to get file from request", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}
	defer file.Close()

	bookIDStr := c.PostForm("book_id")
	if bookIDStr == "" {
		logger.Error("Book ID is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book ID is required"})
		return
	}

	bookID, err := strconv.ParseInt(bookIDStr, 10, 64)
	if err != nil {
		logger.Error("Invalid book ID", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	format := "unknown"
	contentType := header.Header.Get("Content-Type")
	if contentType != "" {
		exts, _ := mime.ExtensionsByType(contentType)
		if len(exts) > 0 {
			format = strings.TrimPrefix(exts[0], ".")
		}
	}

	if format == "unknown" {
		ext := filepath.Ext(header.Filename)
		if ext != "" {
			format = strings.TrimPrefix(ext, ".")
		}
	}

	meta := library_models.FileMetadata{
		BookID: bookID,
		Format: format,
	}

	uploadedFile, err := h.storage.SaveFile(c.Request.Context(), meta, file)
	if err != nil {
		logger.Error("Failed to save file", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	upperCaseFormat := strings.ToUpper(format)

	createReq := &library_models.CreateAttachmentRequest{
		BookId: bookID,
		Format: upperCaseFormat,
		FileId: uploadedFile.FilePath,
	}

	attachment, err := h.service.CreateAttachment(c.Request.Context(), createReq)
	if err != nil {
		if delErr := h.storage.DeleteFile(c.Request.Context(), uploadedFile.FilePath); delErr != nil {
			logger.Error("Failed to delete file after DB error", "file_path", uploadedFile.FilePath, "error", delErr)
		}

		logger.Error("Failed to create attachment record", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create attachment record"})
		return
	}

	logger.Info("File uploaded and attachment created successfully", "book_id", bookID, "format", format, "file_id", uploadedFile.FilePath)
	c.JSON(http.StatusCreated, library_models.AttachmentFromProto(attachment))
}
