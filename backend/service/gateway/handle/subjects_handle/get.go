package subjects_handle

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) get(c *gin.Context) {
	const op = "gateway.Subjects.Get"
	logger := h.logger.With(slog.String("op", op))

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error("Invalid subject ID", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subject ID"})
		return
	}

	res, err := h.service.GetSubject(c.Request.Context(), int32(id))
	if err != nil {
		logger.Error("Failed to get subject", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get subject"})
		return
	}

	logger.Info("Subject retrieved successfully")
	c.JSON(http.StatusOK, res)
}
