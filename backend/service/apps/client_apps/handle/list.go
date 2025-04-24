package handle

import (
	pb "backend/protos/gen/go/apps/clients_apps"
	"backend/service/apps/models"
	"backend/service/utils"
	"context"
	"log/slog"
)

func (s *serverAPI) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	const op = "grpc.handler.ClientApp.List"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("starting operation", slog.Int("page", int(req.GetPage())), slog.Int("count", int(req.GetCount())))

	if err := utils.ValidatePagination(int(req.GetPage()), int(req.GetCount())); err != nil {
		logger.Warn("pagination validation failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	filter := models.ListFilter{
		Page:  int(req.GetPage()),
		Count: int(req.GetCount()),
	}

	if req.ClientId != nil {
		if err := utils.ValidateUUIDToString(*req.ClientId); err != nil {
			logger.Warn("client_id validation failed", slog.Any("error", err))
			return nil, s.convertError(err)
		}
		filter.ClientID = req.ClientId
	}

	if req.AppId != nil {
		appID := int(*req.AppId)
		if err := utils.ValidateAppID(appID); err != nil {
			logger.Warn("app_id validation failed", slog.Any("error", err))
			return nil, s.convertError(err)
		}
		filter.AppID = &appID
	}

	if req.IsActive != nil {
		filter.IsActive = req.IsActive
	}

	clientApps, total, err := s.service.List(ctx, filter)
	if err != nil {
		return nil, s.convertError(err)
	}

	response := &pb.ListResponse{
		ClientApps: s.convertToPbClientApps(clientApps),
		TotalCount: int32(total),
		Page:       req.GetPage(),
		Count:      req.GetCount(),
	}

	logger.Info("operation completed successfully", slog.Int("returned_items", len(clientApps)), slog.Int("total_count", int(total)))
	return response, nil
}
