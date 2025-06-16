package teachers_handle

import (
	sl "backend/pkg/logger"
	"context"
	"log/slog"

	"backend/pkg/models/private_school"
	"backend/protos/gen/go/private_school"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) ListTeachers(ctx context.Context, req *private_school_v1.ListTeachersRequest) (*private_school_v1.ListTeachersResponse, error) {
	const op = "grpc.PrivateSchool.Teachers.ListTeachers"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", req.GetClientId()))
	logger.Debug("List teachers called")

	filter, err := private_school_models.ListTeachersFilterFromProto(req)
	if err != nil {
		logger.Warn("Invalid list parameters", sl.Err(err, true))
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if err = filter.Validate(); err != nil {
		logger.Warn("Validation failed", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	response, err := s.service.ListTeachers(ctx, filter)
	if err != nil {
		logger.Error("Failed to list teachers", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	protoResp := private_school_models.ListTeachersResponseToProto(response)
	logger.Info("Teachers listed successfully", slog.Int("count", len(protoResp.Teachers)), slog.Bool("has_next", protoResp.NextCursor != nil))
	return protoResp, nil
}
