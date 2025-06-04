package subjects_handle

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) list(c *gin.Context) {
	const op = "gateway.Subjects.List"
	logger := h.logger.With(slog.String("op", op))

	res, err := h.service.ListSubjects(c.Request.Context())
	if err != nil {
		logger.Error("Failed to list subjects", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list subjects"})
		return
	}

	logger.Info("Subjects listed successfully")
	c.JSON(http.StatusOK, res)
}
