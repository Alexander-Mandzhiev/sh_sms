package students_handle

import (
	"backend/pkg/models/students"
	"backend/protos/gen/go/private_school"
	"context"
	"log/slog"
)

func (s *serverAPI) UpdateStudent(ctx context.Context, req *private_school_v1.UpdateStudentRequest) (*private_school_v1.StudentResponse, error) {
	const op = "grpc.PrivateSchool.StudentService.UpdateStudent"
	logger := s.logger.With(slog.String("op", op), slog.String("student_id", req.GetId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("UpdateStudent called")

	updateData, err := students_models.UpdateStudentFromProto(req)
	if err != nil {
		logger.Warn("Invalid update parameters", slog.String("error", err.Error()))
		return nil, s.convertError(err)
	}

	updatedStudent, err := s.service.UpdateStudent(ctx, updateData)
	if err != nil {
		logger.Error("Failed to update student", slog.String("error", err.Error()))
		return nil, s.convertError(err)
	}

	response := updatedStudent.StudentToProto()
	logger.Info("Student updated successfully")
	return response, nil
}
