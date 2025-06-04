package books_handle

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) delete(c *gin.Context) {
	const op = "gateway.Books.Delete"
	logger := h.logger.With(slog.String("op", op))

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error("Invalid book ID", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	clientID := c.Query("client_id")
	if clientID == "" {
		logger.Error("Client ID is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Client ID is required"})
		return
	}

	if err = h.service.DeleteBook(c.Request.Context(), id, clientID); err != nil {
		logger.Error("Failed to delete book", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	logger.Info("Book deleted successfully")
	c.Status(http.StatusNoContent)
}
