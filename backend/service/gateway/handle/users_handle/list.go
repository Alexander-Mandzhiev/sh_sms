package users_handle

import (
	"log/slog"
	"net/http"

	"backend/protos/gen/go/sso/users"
	"backend/service/gateway/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) listUsers(c *gin.Context) {
	const op = "gateway.Users.List"
	logger := h.logger.With(slog.String("op", op))

	var req models.ListUsersRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Invalid request format", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	grpcReq := &users.ListRequest{
		ClientId: req.ClientID,
		Page:     int32(req.Pagination.Page),
		Count:    int32(req.Pagination.Count),
	}

	if req.Filters.Email != nil {
		grpcReq.EmailFilter = req.Filters.Email
	}
	if req.Filters.Phone != nil {
		grpcReq.PhoneFilter = req.Filters.Phone
	}
	if req.Filters.ActiveOnly != nil {
		grpcReq.ActiveOnly = req.Filters.ActiveOnly
	}

	resp, err := h.service.ListUsers(c.Request.Context(), grpcReq)
	if err != nil {
		logger.Error("Failed to list users", slog.String("client_id", req.ClientID), slog.String("error", err.Error()))
		h.handleGRPCError(c, err)
		return
	}

	usersList := make([]*models.User, 0, len(resp.Users))
	for _, protoUser := range resp.Users {
		user, err := models.UserFromProto(protoUser)
		if err != nil {
			logger.Error("Proto conversion failed",
				slog.String("error", err.Error()))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		usersList = append(usersList, user)
	}

	logger.Info("Users listed successfully", slog.Int("count", len(usersList)), slog.Int("total", int(resp.TotalCount)))
	c.JSON(http.StatusOK, gin.H{
		"data": usersList,
		"meta": gin.H{
			"total_count":    resp.TotalCount,
			"current_page":   resp.CurrentPage,
			"count_per_page": req.Pagination.Count,
		},
	})
}
