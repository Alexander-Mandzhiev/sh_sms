package students_handle

import (
	"backend/pkg/models/students"
	"backend/pkg/utils"
	"backend/protos/gen/go/private_school"
	"context"
	"log/slog"
)

func (s *serverAPI) GetStudent(ctx context.Context, req *private_school_v1.StudentRequest) (*private_school_v1.StudentResponse, error) {
	const op = "grpc.StudentService.GetStudent"
	logger := s.logger.With(slog.String("op", op), slog.String("student_id", req.GetId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("GetStudent called")

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

	student, err := s.service.GetStudent(ctx, studentID, clientID)
	if err != nil {
		logger.Error("Failed to get student", slog.String("error", err.Error()))
		return nil, s.convertError(err)
	}

	// Проверка на soft-delete (если требуется)
	if !student.IsActive() {
		logger.Warn("Student is soft-deleted", slog.String("student_id", studentID.String()))
		return nil, s.convertError(students_models.ErrStudentNotFound)
	}

	// Преобразование результата в protobuf
	response := student.StudentToProto()
	logger.Debug("Student retrieved successfully")
	return response, nil
}
