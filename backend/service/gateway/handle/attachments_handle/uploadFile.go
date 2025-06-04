package attachments_handle

import (
	library "backend/protos/gen/go/library"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) upload(c *gin.Context) {
	const op = "gateway.Attachments.Upload"
	logger := h.logger.With(slog.String("op", op))

	if err := c.Request.ParseMultipartForm(50 << 20); err != nil {
		logger.Error("Failed to parse multipart form", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
		return
	}

	bookID, err := strconv.ParseInt(c.PostForm("book_id"), 10, 64)
	if err != nil {
		logger.Error("Invalid book ID", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	format := c.PostForm("format")
	if format == "" {
		logger.Error("Format parameter is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format parameter is required"})
		return
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		logger.Error("Failed to get file", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		logger.Error("Failed to open file", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process file"})
		return
	}
	defer file.Close()

	res, err := h.service.UploadFile(c.Request.Context(), &library.FileMetadata{BookId: bookID, Format: format}, file)
	if err != nil {
		logger.Error("Failed to upload file", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	logger.Info("File uploaded successfully")
	c.JSON(http.StatusOK, res)
}
