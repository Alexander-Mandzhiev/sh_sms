package roles_handle

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/sso/roles"
	"backend/service/gateway/models"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) createRole(c *gin.Context) {
	const op = "gateway.Roles.Create"
	logger := h.logger.With(slog.String("op", op))

	var req models.CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if req.ClientID == "" || req.AppID == 0 || req.Name == "" || req.Level == 0 {
		logger.Warn("Missing required fields", "client_id", req.ClientID, "app_id", req.AppID, "name", req.Name, "level", req.Level)
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required fields"})
		return
	}

	grpcReq := &roles.CreateRequest{
		ClientId:    req.ClientID,
		AppId:       int32(req.AppID),
		Name:        req.Name,
		Description: req.Description,
		Level:       int32(req.Level),
		IsCustom:    req.IsCustom,
		CreatedBy:   utils.StringPtr(req.CreatedBy.String()),
	}

	resp, err := h.service.CreateRole(c.Request.Context(), grpcReq)
	if err != nil {
		logger.Error("Failed to create role", "error", err)
		h.handleGRPCError(c, err)
		return
	}

	role, err := models.RoleFromProto(resp)
	if err != nil {
		logger.Error("Failed to convert proto role", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	logger.Info("Role created successfully", "role_id", role.ID, "client_id", role.ClientID, "app_id", role.AppID)
	c.JSON(http.StatusCreated, role)
}
