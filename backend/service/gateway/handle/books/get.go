package books_handle

import (
	library_models "backend/pkg/models/library"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) get(c *gin.Context) {
	const op = "gateway.Books.Get"
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

	res, err := h.service.GetBook(c.Request.Context(), id, clientID)
	if err != nil {
		logger.Error("Failed to get book", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get book"})
		return
	}

	response := library_models.BookResponseFromProto(res)
	logger.Info("Book retrieved successfully")
	c.JSON(http.StatusOK, response)
}
