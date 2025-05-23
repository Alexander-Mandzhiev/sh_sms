package auth_handle

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/auth"
	"backend/service/gateway/models/auth"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) listAllSessions(c *gin.Context) {
	const op = "grpc.handler.ListAllSessions"
	logger := h.logger.With(slog.String("op", op))

	var req auth_models.AllSessionsFilter
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("failed to bind request", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := utils.ValidateUUID(req.ClientID); err != nil {
		logger.Error("client_id is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "client_id is required"})
		return
	}

	if err := utils.ValidateAppID(req.AppID); err != nil {
		logger.Error("invalid app_id", slog.Int("app_id", req.AppID))
		c.JSON(http.StatusBadRequest, gin.H{"error": "app_id cannot be negative"})
		return
	}

	if err := utils.ValidatePagination(req.Page, req.Count); err != nil {
		logger.Error("invalid pagination", slog.Int("count", req.Count))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid pagination"})
		return
	}

	pbFilter := &auth.AllSessionsFilter{
		ClientId:   req.ClientID.String(),
		AppId:      int32(req.AppID),
		Page:       int32(req.Page),
		Count:      int32(req.Count),
		ActiveOnly: req.ActiveOnly,
		FullName:   req.FullName,
		Phone:      req.Phone,
		Email:      req.Email,
	}

	sessions, err := h.service.ListAllSessions(c.Request.Context(), pbFilter)
	if err != nil {
		logger.Error("failed to get sessions", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	response := auth_models.SessionsFromProto(sessions.Sessions)
	logger.Info("successfully retrieved sessions", slog.Int("count", len(sessions.Sessions)), slog.Int("page", req.Page))
	c.JSON(http.StatusOK, gin.H{"sessions": response, "count": len(response)})

}
