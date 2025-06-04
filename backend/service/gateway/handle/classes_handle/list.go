package classes_handle

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) list(c *gin.Context) {
	const op = "gateway.Permissions.Delete"
	logger := h.logger.With(slog.String("op", op))

	res, err := h.service.ListClasses(c.Request.Context())
	if err != nil {
		logger.Error("Failed to list classes", slog.String("error", err.Error()))
		c.Status(http.StatusInternalServerError)
		return
	}

	logger.Info("Classes listed successfully")
	c.JSON(http.StatusOK, res)
}
