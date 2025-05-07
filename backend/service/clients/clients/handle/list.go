package handle

import (
	pb "backend/protos/gen/go/clients/clients"
	"backend/service/clients/clients/models"
	"backend/service/utils"
	"context"
	"fmt"
	"log/slog"
)

func (s *serverAPI) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	const op = "grpc.clients.List"
	logger := s.logger.With(slog.String("op", op), slog.Int("page", int(req.GetPage())), slog.Int("count", int(req.GetCount())))
	logger.Debug("processing list request")

	if err := utils.ValidatePagination(int(req.GetPage()), int(req.GetCount())); err != nil {
		logger.Warn("invalid pagination", slog.String("error", err.Error()))
		return nil, s.convertError(fmt.Errorf("%w: pagination", ErrInvalidArgument))
	}

	filter := models.Filter{
		Search:     req.Search,
		ActiveOnly: req.ActiveOnly,
		TypeID:     nil,
	}

	if req.TypeId != nil {
		tid := int(req.GetTypeId())
		if tid <= 0 {
			logger.Warn("invalid type_id", slog.Int("value", tid))
			return nil, s.convertError(fmt.Errorf("%w: type_id", ErrInvalidArgument))
		}
		filter.TypeID = &tid
	}

	pagination := models.Pagination{
		Page:  int(req.GetPage()),
		Count: int(req.GetCount()),
	}

	clients, total, err := s.service.List(ctx, filter, pagination)
	if err != nil {
		logger.Error("list failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	protoClients := make([]*pb.Client, 0, len(clients))
	for _, c := range clients {
		protoClients = append(protoClients, c.ToProto())
	}

	logger.Info("list completed", slog.Int("total", total), slog.Int("returned", len(protoClients)))
	return &pb.ListResponse{
		Clients:    protoClients,
		TotalCount: int32(total),
	}, nil
}
