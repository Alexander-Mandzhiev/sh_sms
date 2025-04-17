package handle

import (
	sl "backend/pkg/logger"
	pb "backend/protos/gen/go/apps/app_manager"
	"backend/service/apps/constants"
	"backend/service/apps/models"
	"backend/service/utils"
	"context"
	"fmt"
	"log/slog"
)

func (s *serverAPI) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	const op = "grpc.handler.AppManager.List"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("List request received", slog.Int64("page", req.GetPage()), slog.Int64("count", req.GetCount()), slog.Any("filter_is_active", req.GetFilterIsActive()))

	if err := utils.ValidatePagination(int(req.GetPage()), int(req.GetCount())); err != nil {
		err = fmt.Errorf("%w: page=%d count=%d", constants.ErrInvalidPagination, req.GetPage(), req.GetCount())
		logger.Warn("Invalid pagination parameters", sl.Err(err, false), slog.Int64("page", req.GetPage()), slog.Int64("count", req.GetCount()))
		return nil, s.convertError(err)
	}

	filter := models.ListFilter{
		Page:     int(req.GetPage()),
		Count:    int(req.GetCount()),
		IsActive: req.FilterIsActive,
	}

	apps, total, err := s.service.List(ctx, filter)
	if err != nil {
		logger.Error("Failed to list apps", sl.Err(err, true), slog.Any("filter", filter))
		return nil, s.convertError(err)
	}

	pbApps := make([]*pb.App, 0, len(apps))
	for _, app := range apps {
		pbApps = append(pbApps, s.convertAppToProto(&app))
	}

	logger.Info("List completed successfully", slog.Int("returned_items", len(pbApps)), slog.Int("total_items", total))

	return &pb.ListResponse{
		Apps:       pbApps,
		TotalCount: int32(total),
		Page:       req.Page,
		Count:      req.Count,
	}, nil
}
