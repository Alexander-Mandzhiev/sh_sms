package books_handle

import (
	"backend/pkg/models/library"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) update(c *gin.Context) {
	const op = "gateway.Books.Update"
	logger := h.logger.With(slog.String("op", op))

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error("Invalid book ID", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var params library_models.UpdateBookRequest
	if err = c.ShouldBindJSON(&params); err != nil {
		logger.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	params.ID = id
	res, err := h.service.UpdateBook(c.Request.Context(), &params)
	if err != nil {
		logger.Error("Failed to update book", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	response := library_models.BookResponseFromProto(res)
	logger.Info("Book updated successfully")
	c.JSON(http.StatusOK, response)
}
