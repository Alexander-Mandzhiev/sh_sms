package handle

import (
	"backend/protos/gen/go/sso/roles"
	"backend/service/sso/models"
	"backend/service/utils"
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *serverAPI) Create(ctx context.Context, req *roles.CreateRequest) (*roles.Role, error) {
	const op = "grpc.role.Create"
	logger := s.logger.With(slog.String("op", op), slog.String("role_name", req.GetName()), slog.String("client_id", req.GetClientId()))
	logger.Debug("attempting to create role")

	clientID, err := utils.ValidateAndReturnUUID(req.GetClientId())
	if err != nil {
		logger.Warn("invalid client_id", slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: client_id", ErrInvalidArgument))
	}

	if err = utils.ValidateRoleName(req.GetName()); err != nil {
		logger.Warn("invalid role name", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	if err = utils.ValidateRoleLevel(req.GetLevel()); err != nil {
		logger.Warn("invalid role level", slog.Int("level", int(req.GetLevel())), slog.Any("error", err))
		return nil, s.convertError(err)
	}

	role := &models.Role{
		ClientID:    clientID,
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Level:       int(req.GetLevel()),
		IsCustom:    req.GetIsCustom(),
	}

	if req.GetCreatedBy() != "" {
		var createdBy uuid.UUID
		createdBy, err = utils.ValidateAndReturnUUID(req.GetCreatedBy())
		if err != nil {
			logger.Warn("invalid created_by", slog.Any("error", err))
			return nil, s.convertError(fmt.Errorf("%w: created_by", ErrInvalidArgument))
		}
		role.CreatedBy = &createdBy
	}

	if err = s.service.Create(ctx, role); err != nil {
		logger.Error("create role failed", slog.Any("error", err), slog.Any("role", role))
		return nil, s.convertError(err)
	}

	logger.Info("role created successfully", slog.String("role_id", role.ID.String()))
	return convertRoleToProto(role), nil
}
