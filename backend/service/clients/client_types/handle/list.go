package handle

import (
	"backend/protos/gen/go/clients/client_types"
	"backend/service/clients/client_types/models"
	"backend/service/utils"
	"context"
	"fmt"
	"log/slog"
)

func (s *serverAPI) ListClientType(ctx context.Context, req *client_types.ListRequest) (*client_types.ListResponse, error) {
	const op = "grpc.client_types.List"
	logger := s.logger.With(slog.String("op", op), slog.Int("page", int(req.GetPage())), slog.Int("page_size", int(req.GetCount())))
	logger.Debug("processing list client types request")

	if err := utils.ValidatePagination(int(req.GetPage()), int(req.GetCount())); err != nil {
		err = fmt.Errorf("%w: page must be ≥ 1", ErrInvalidArgument)
		logger.Warn("invalid page number", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	filter := models.Filter{
		Search:     req.Search,
		ActiveOnly: req.ActiveOnly,
	}

	pagination := models.Pagination{
		Page:     int(req.GetPage()),
		PageSize: int(req.GetCount()),
	}

	clientTypes, total, err := s.service.List(ctx, filter, pagination)
	if err != nil {
		logger.Error("failed to list client types", slog.String("error_type", "service_error"), slog.Any("error", err))
		return nil, s.convertError(err)
	}

	pbTypes := make([]*client_types.ClientType, 0, len(clientTypes))
	for _, ct := range clientTypes {
		pbTypes = append(pbTypes, models.ToProto(ct))
	}

	logger.Info("client types listed successfully", slog.Int("returned_items", len(pbTypes)), slog.Int("total_items", total))
	return &client_types.ListResponse{
		ClientTypes: pbTypes,
		TotalCount:  int32(total),
		CurrentPage: req.GetPage(),
	}, nil
}
