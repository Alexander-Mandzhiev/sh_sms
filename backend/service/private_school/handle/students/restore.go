package students_handle

import (
	"backend/pkg/models/students"
	"backend/pkg/utils"
	"backend/protos/gen/go/private_school"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

func (s *serverAPI) RestoreStudent(ctx context.Context, req *private_school_v1.StudentRequest) (*emptypb.Empty, error) {
	const op = "grpc.StudentService.RestoreStudent"
	logger := s.logger.With(slog.String("op", op), slog.String("student_id", req.GetId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("RestoreStudent called")

	studentID, err := utils.ValidateStringAndReturnUUID(req.GetId())
	if err != nil {
		logger.Warn("Invalid student ID format", slog.String("error", err.Error()))
		return nil, s.convertError(students_models.ErrInvalidStudentID)
	}

	clientID, err := utils.ValidateStringAndReturnUUID(req.GetClientId())
	if err != nil {
		logger.Warn("Invalid client ID format", slog.String("error", err.Error()))
		return nil, s.convertError(students_models.ErrInvalidClientID)
	}

	if err = s.service.RestoreStudent(ctx, studentID, clientID); err != nil {
		logger.Error("Failed to restore student", slog.String("error", err.Error()))
		return nil, s.convertError(err)
	}

	logger.Info("Student restored successfully")
	return &emptypb.Empty{}, nil
}
