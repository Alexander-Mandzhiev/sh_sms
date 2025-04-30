package handle

import (
	"backend/protos/gen/go/clients/client_types"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

func (s *serverAPI) Delete(ctx context.Context, req *client_types.DeleteRequest) (*emptypb.Empty, error) {
	const op = "grpc.client_types.Delete"
	logger := s.logger.With(slog.String("op", op), slog.Int("id", int(req.GetId())), slog.Bool("permanent", req.GetPermanent()))
	logger.Debug("processing client type deletion")

	if req.GetId() <= 0 {
		err := fmt.Errorf("%w: invalid client type ID", ErrInvalidArgument)
		logger.Warn("validation failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	err := s.service.Delete(ctx, int(req.GetId()), req.GetPermanent())
	if err != nil {
		logger.Error("client type deletion failed", slog.String("error_type", "service_error"), slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("client type deleted successfully", slog.Bool("permanent", req.GetPermanent()))
	return &emptypb.Empty{}, nil
}
