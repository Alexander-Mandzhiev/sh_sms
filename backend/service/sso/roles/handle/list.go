package handle

import (
	"backend/protos/gen/go/sso/roles"
	"backend/service/constants"
	"backend/service/sso/models"
	"backend/service/utils"
	"context"
	"fmt"
	"log/slog"
)

func (s *serverAPI) List(ctx context.Context, req *roles.ListRequest) (*roles.ListResponse, error) {
	const op = "grpc.roles.List"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", req.GetClientId()))
	logger.Debug("attempting to list roles")

	clientID, err := utils.ValidateAndReturnUUID(req.GetClientId())
	if err != nil {
		logger.Warn("invalid client_id", slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: client_id", constants.ErrInvalidArgument))
	}

	listReq := models.ListRequest{
		ClientID:    &clientID,
		NameFilter:  req.NameFilter,
		LevelFilter: utils.GetIntPointer(int(req.GetLevelFilter())),
		ActiveOnly:  req.ActiveOnly,
		Page:        int(req.GetPage()),
		Count:       int(req.GetCount()),
	}

	if err = utils.ValidatePagination(listReq.Page, listReq.Count); err != nil {
		logger.Warn("invalid pagination", slog.Int("page", listReq.Page), slog.Int("count", listReq.Count), slog.Any("error", err))
		return nil, s.convertError(err)
	}

	rolesList, total, err := s.service.List(ctx, listReq)
	if err != nil {
		logger.Error("list roles failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	protoRoles := make([]*roles.Role, 0, len(rolesList))
	for _, r := range rolesList {
		protoRoles = append(protoRoles, convertRoleToProto(&r))
	}

	logger.Debug("list roles successful", slog.Int("count", len(protoRoles)), slog.Int("total", total))
	return &roles.ListResponse{
		Roles:       protoRoles,
		TotalCount:  int32(total),
		CurrentPage: int32(listReq.Page),
	}, nil
}
