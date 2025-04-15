package handle

import (
	sl "backend/pkg/logger"
	pb "backend/protos/gen/go/apps/app_manager"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) Delete(ctx context.Context, req *pb.AppIdentifier) (*pb.DeleteResponse, error) {
	const op = "grpc.handler.Delete"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("Delete request received", slog.Any("request", req))

	if err := validateNoConflict(req); err != nil {
		logger.Warn("Conflicting identifiers", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	if err := validateAtLeastOne(req); err != nil {
		logger.Warn("Missing identifier", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	if id := req.GetId(); id != 0 {
		if valErr := validateID(id); valErr != nil {
			logger.Warn("Invalid ID format", sl.Err(valErr, true), slog.Int("requested_id", int(id)))
			return nil, s.convertError(valErr)
		}

		if err := s.service.DeleteByID(ctx, int(id)); err != nil {
			logger.Error("Failed to delete by ID", sl.Err(err, true), slog.Int("converted_id", int(id)))
			return &pb.DeleteResponse{Success: false}, s.convertError(err)
		}

		logger.Info("Successfully deleted app by ID", slog.Int("deleted_id", int(id)))
		return &pb.DeleteResponse{Success: true}, nil
	}

	if code := req.GetCode(); code != "" {
		if valErr := validateCode(code, 50); valErr != nil {
			logger.Warn("Invalid code format", sl.Err(valErr, true), slog.String("requested_code", code))
			return nil, s.convertError(valErr)
		}

		if err := s.service.DeleteByCode(ctx, code); err != nil {
			logger.Error("Failed to delete by code", sl.Err(err, true), slog.String("requested_code", code))
			return &pb.DeleteResponse{Success: false}, s.convertError(err)
		}

		logger.Info("Successfully deleted app by code", slog.String("deleted_code", code))
		return &pb.DeleteResponse{Success: true}, nil
	}

	logger.Error("Unexpected state in Delete handler")
	return nil, status.Error(codes.Internal, "internal server error")
}
