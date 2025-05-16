package handle

import (
	utils2 "backend/pkg/utils"
	"backend/protos/gen/go/sso/roles"
	"backend/service/constants"
	"backend/service/sso/models"
	"context"
	"fmt"
	"log/slog"
)

func (s *serverAPI) ListRoles(ctx context.Context, req *roles.ListRequest) (*roles.ListResponse, error) {
	const op = "grpc.roles.List"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", req.GetClientId()))
	logger.Debug("attempting to list roles")

	clientID, err := utils2.ValidateAndReturnUUID(req.GetClientId())
	if err != nil {
		logger.Warn("invalid client_id", slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: client_id", constants.ErrInvalidArgument))
	}
	if err = utils2.ValidatePagination(int(req.GetPage()), int(req.GetCount())); err != nil {
		logger.Warn("invalid pagination", slog.Int("page", int(req.GetPage())), slog.Int("count", int(req.GetCount())), slog.Any("error", err))
		return nil, s.convertError(err)
	}

	if err = utils2.ValidateAppID(int(req.GetAppId())); err != nil {
		logger.Warn("invalid app_id", slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: app_id", constants.ErrInvalidArgument))
	}

	listReq := models.ListRequest{
		ClientID: &clientID,
		AppID:    utils2.GetIntPointer(int(req.GetAppId())),
		Page:     int(req.GetPage()),
		Count:    int(req.GetCount()),
	}

	if req.LevelFilter != nil {
		listReq.LevelFilter = utils2.GetIntPointer(int(req.GetLevelFilter()))
	}

	if req.GetNameFilter() != "" {
		nameFilter := req.GetNameFilter()
		listReq.NameFilter = &nameFilter
	}

	if req.ActiveOnly != nil {
		listReq.ActiveOnly = req.ActiveOnly
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
