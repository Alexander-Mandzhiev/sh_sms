package subjects_handle

import (
	library_models "backend/pkg/models/library"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) create(c *gin.Context) {
	const op = "gateway.Subjects.Create"
	logger := h.logger.With(slog.String("op", op))

	var params library_models.CreateSubjectParams
	if err := c.ShouldBindJSON(&params); err != nil {
		logger.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	res, err := h.service.CreateSubject(c.Request.Context(), &params)
	if err != nil {
		logger.Error("Failed to create subject", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subject"})
		return
	}

	logger.Info("Subject created successfully")
	c.JSON(http.StatusCreated, res)
}
