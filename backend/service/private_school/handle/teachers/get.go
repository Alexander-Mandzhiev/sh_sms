package teachers_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/private_school"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) GetTeacher(ctx context.Context, req *private_school_v1.GetTeacherRequest) (*private_school_v1.TeacherResponse, error) {
	const op = "grpc.PrivateSchool.Teachers.GetTeacher"
	logger := s.logger.With(slog.String("op", op), slog.String("teacher_id", req.GetId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("Get teacher called")

	id, clientID, err := validateUUIDs(req.GetId(), req.GetClientId())
	if err != nil {
		logger.Error("Validation failed", sl.Err(err, true))
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	teacher, err := s.service.GetTeacher(ctx, id, clientID)
	if err != nil {
		logger.Error("Failed to get teacher", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	proto := teacher.TeacherToProto()
	logger.Info("Teacher retrieved successfully")
	return proto, nil
}
