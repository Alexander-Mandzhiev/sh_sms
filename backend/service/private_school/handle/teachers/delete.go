package teachers_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/private_school"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

func (s *serverAPI) DeleteTeacher(ctx context.Context, req *private_school_v1.DeleteTeacherRequest) (*emptypb.Empty, error) {
	const op = "grpc.PrivateSchool.Teachers.DeleteTeacher"
	logger := s.logger.With(slog.String("op", op), slog.String("teacher_id", req.GetId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("Delete teacher called")

	id, clientID, err := validateUUIDs(req.GetId(), req.GetClientId())
	if err != nil {
		logger.Error("Validation failed", sl.Err(err, true))
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if err = s.service.DeleteTeacher(ctx, id, clientID); err != nil {
		logger.Error("Failed to get teacher", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	logger.Info("Teacher deleted successfully")
	return &emptypb.Empty{}, nil
}
