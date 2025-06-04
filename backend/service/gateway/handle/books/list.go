package books_handle

import (
	library_models "backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) list(c *gin.Context) {
	const op = "gateway.Books.List"
	logger := h.logger.With(slog.String("op", op))

	clientID := c.Query("client_id")
	if clientID == "" {
		logger.Error("Client ID is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Client ID is required"})
		return
	}

	count, _ := strconv.Atoi(c.Query("count"))
	cursor, _ := strconv.ParseInt(c.Query("cursor"), 10, 64)
	filter := c.Query("filter")

	req := &library.ListBooksRequest{
		ClientId: clientID,
		Filter:   &filter,
	}

	if count > 0 {
		count32 := int32(count)
		req.Count = &count32
	}

	if cursor > 0 {
		req.Cursor = &cursor
	}

	res, err := h.service.ListBooks(c.Request.Context(), req)
	if err != nil {
		logger.Error("Failed to list books", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list books"})
		return
	}

	response := library_models.ListBooksResponseFromProto(res)
	logger.Info("Books listed successfully", "count", len(res.Books))
	c.JSON(http.StatusOK, response)
}
