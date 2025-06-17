package teachers_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/models/teacher"
	"backend/protos/gen/go/private_school"
	"context"
	"log/slog"
)

func (s *serverAPI) CreateTeacher(ctx context.Context, req *private_school_v1.CreateTeacherRequest) (*private_school_v1.TeacherResponse, error) {
	const op = "grpc.PrivateSchool.Teachers.CreateTeacher"
	logger := s.logger.With(slog.String("op", op), slog.String("teacher_id", req.GetId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("Create teacher called")

	createTeacher, err := teachers_models.CreateTeacherFromProto(req)
	if err != nil {
		logger.Warn("Invalid create subject parameters", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	teacher, err := s.service.CreateTeacher(ctx, createTeacher)
	if err != nil {
		logger.Error("Failed to create subject", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	proto := teacher.TeacherToProto()
	logger.Info("Teacher created successfully")
	return proto, nil
}
