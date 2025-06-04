package subjects_handle

import (
	library_models "backend/pkg/models/library"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) update(c *gin.Context) {
	const op = "gateway.Subjects.Update"
	logger := h.logger.With(slog.String("op", op))

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error("Invalid subject ID", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subject ID"})
		return
	}

	var params library_models.UpdateSubjectParams
	if err := c.ShouldBindJSON(&params); err != nil {
		logger.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	params.ID = int32(id)
	res, err := h.service.UpdateSubject(c.Request.Context(), &params)
	if err != nil {
		logger.Error("Failed to update subject", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update subject"})
		return
	}

	logger.Info("Subject updated successfully")
	c.JSON(http.StatusOK, res)
}
