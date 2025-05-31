package classes_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

func (s *serverAPI) ListClasses(ctx context.Context, _ *emptypb.Empty) (*library.ListClassesResponse, error) {
	const op = "grpc.Library.Classes.ListClasses"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("List classes called")

	classes, err := s.service.List(ctx)
	if err != nil {
		logger.Error("Failed to list classes", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	protoClasses := make([]*library.Class, 0, len(classes))
	for _, class := range classes {
		protoClasses = append(protoClasses, library_models.ClassToProto(class))
	}

	resp := &library.ListClassesResponse{
		Classes: protoClasses,
	}

	logger.Info("Classes listed", slog.Int("count", len(protoClasses)))
	return resp, nil
}
