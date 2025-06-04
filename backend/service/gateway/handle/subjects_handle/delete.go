package subjects_handle

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) delete(c *gin.Context) {
	const op = "gateway.Subjects.Delete"
	logger := h.logger.With(slog.String("op", op))

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error("Invalid subject ID", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subject ID"})
		return
	}

	err = h.service.DeleteSubject(c.Request.Context(), int32(id))
	if err != nil {
		logger.Error("Failed to delete subject", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete subject"})
		return
	}

	logger.Info("Subject deleted successfully")
	c.Status(http.StatusNoContent)
}
