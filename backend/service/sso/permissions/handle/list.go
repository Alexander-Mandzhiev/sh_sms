package handle

import (
	"backend/protos/gen/go/sso/permissions"
	"backend/service/sso/models"
	"backend/service/utils"
	"context"
	"log/slog"
)

func (s *serverAPI) List(ctx context.Context, req *permissions.ListRequest) (*permissions.ListResponse, error) {
	const op = "grpc.Permission.List"
	logger := s.logger.With(slog.String("op", op), slog.Int("app_id", int(req.GetAppId())))
	logger.Debug("processing list request")

	if err := utils.ValidateAppID(int(req.GetAppId())); err != nil {
		logger.Warn("invalid app_id", slog.Any("error", err))
		return nil, s.convertError(ErrInvalidArgument)
	}

	if err := utils.ValidatePagination(int(req.GetPage()), int(req.GetCount())); err != nil {
		logger.Warn("invalid pagination", slog.Int("page", int(req.GetPage())), slog.Int("count", int(req.GetCount())), slog.Any("error", err))
		return nil, s.convertError(err)
	}

	filter := models.ListRequest{
		AppID: utils.GetIntPointer(int(req.GetAppId())),
		Page:  int(req.GetPage()),
		Count: int(req.GetCount()),
	}

	if req.ActiveOnly != nil {
		filter.ActiveOnly = req.ActiveOnly
	}
	if req.GetCodeFilter() != "" {
		filter.CodeFilter = req.CodeFilter
	}
	if req.GetCategory() != "" {
		filter.CategoryFilter = req.Category
	}

	perms, total, err := s.service.List(ctx, filter)
	if err != nil {
		logger.Error("list permissions failed", slog.Any("error", err), slog.Int("app_id", int(req.GetAppId())))
		return nil, s.convertError(err)
	}

	resp := &permissions.ListResponse{
		Permissions: convertPermissionsToProto(perms),
		TotalCount:  int32(total),
		CurrentPage: req.GetPage(),
	}

	logger.Debug("list request processed successfully", slog.Int("found", len(perms)), slog.Int("total", total))
	return resp, nil
}
