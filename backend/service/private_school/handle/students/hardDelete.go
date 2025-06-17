package students_handle

import (
	"backend/pkg/models/students"
	"backend/pkg/utils"
	"backend/protos/gen/go/private_school"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

func (s *serverAPI) HardDeleteStudent(ctx context.Context, req *private_school_v1.StudentRequest) (*emptypb.Empty, error) {
	const op = "grpc.PrivateSchool.StudentService.HardDeleteStudent"
	logger := s.logger.With(slog.String("op", op), slog.String("student_id", req.GetId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("HardDeleteStudent called")

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

	if err = s.service.HardDeleteStudent(ctx, studentID, clientID); err != nil {
		logger.Error("Failed to hard delete student", slog.String("error", err.Error()))
		return nil, s.convertError(err)
	}

	logger.Info("Student hard deleted successfully")
	return &emptypb.Empty{}, nil
}
