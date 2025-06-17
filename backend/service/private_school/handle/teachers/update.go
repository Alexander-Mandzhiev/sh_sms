package teachers_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/models/teacher"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"

	"backend/protos/gen/go/private_school"
)

func (s *serverAPI) UpdateTeacher(ctx context.Context, req *private_school_v1.UpdateTeacherRequest) (*private_school_v1.TeacherResponse, error) {
	const op = "grpc.PrivateSchool.Teachers.UpdateTeacher"
	logger := s.logger.With(slog.String("op", op), slog.String("teacher_id", req.GetId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("Update teacher called")

	update, err := teachers_models.UpdateTeacherFromProto(req)
	if err != nil {
		logger.Warn("Invalid update subject parameters", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	if update.FullName == nil && update.Phone == nil && update.Email == nil && update.AdditionalInfo == nil {
		logger.Warn("No fields to update")
		return nil, status.Error(codes.InvalidArgument, "no fields to update")
	}

	teacher, err := s.service.UpdateTeacher(ctx, update)
	if err != nil {
		logger.Error("Failed to update subject", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	proto := teacher.TeacherToProto()
	logger.Info("Subject created successfully")
	return proto, nil
}
