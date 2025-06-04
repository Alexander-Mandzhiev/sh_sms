package classes_handle

import (
	library "backend/protos/gen/go/library"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) get(c *gin.Context) {
	const op = "gateway.Permissions.Delete"
	logger := h.logger.With(slog.String("op", op))

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error("Invalid class ID", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}
	res, err := h.service.GetClass(c.Request.Context(), &library.GetClassRequest{Id: int32(id)})
	if err != nil {
		logger.Error("Error getting class", slog.Int("id", id), slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Info("Class getting successfully")
	c.JSON(http.StatusOK, res)
}
