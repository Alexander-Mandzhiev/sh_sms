package books_handle

import (
	"backend/pkg/models/library"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) create(c *gin.Context) {
	const op = "gateway.Books.Create"
	logger := h.logger.With(slog.String("op", op))

	var params library_models.CreateBookParams
	if err := c.ShouldBindJSON(&params); err != nil {
		logger.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := params.Validate(); err != nil {
		logger.Error("Validation failed", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	params.Sanitize()

	res, err := h.service.CreateBook(c.Request.Context(), &params)
	if err != nil {
		logger.Error("Failed to create book", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	response := library_models.BookResponseFromProto(res)
	logger.Info("Book created successfully")
	c.JSON(http.StatusCreated, response)
}
