package students_handle

import (
	"backend/pkg/models/students"
	"backend/protos/gen/go/private_school"
	"context"
	"log/slog"
)

func (s *serverAPI) CreateStudent(ctx context.Context, req *private_school_v1.CreateStudentRequest) (*private_school_v1.StudentResponse, error) {
	const op = "grpc.StudentService.CreateStudent"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", req.GetClientId()))
	logger.Debug("CreateStudent called")

	student, err := students_models.CreateStudentFromProto(req)
	if err != nil {
		logger.Warn("Invalid create student parameters", slog.String("error", err.Error()))
		return nil, s.convertError(err)
	}

	createdStudent, err := s.service.CreateStudent(ctx, student)
	if err != nil {
		logger.Error("Failed to create student", slog.String("error", err.Error()))
		return nil, s.convertError(err)
	}

	response := createdStudent.StudentToProto()
	logger.Info("Student created successfully", slog.String("student_id", createdStudent.ID.String()))
	return response, nil
}
