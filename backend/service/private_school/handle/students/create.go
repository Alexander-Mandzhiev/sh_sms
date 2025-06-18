package students_handle

import (
	"backend/pkg/models/students"
	"backend/pkg/utils"
	"backend/protos/gen/go/private_school"
	"context"
	"log/slog"
)

func (s *serverAPI) CreateStudent(ctx context.Context, req *private_school_v1.CreateStudentRequest) (*private_school_v1.StudentResponse, error) {
	const op = "grpc.StudentService.CreateStudent"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", req.GetClientId()), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("request received")

	if ctx.Err() != nil {
		logger.Warn("request canceled before processing")
		return nil, s.convertError(ctx.Err())
	}

	student, err := students_models.CreateStudentFromProto(req)
	if err != nil {
		logger.Warn("invalid parameters", "error", err)
		return nil, s.convertError(err)
	}

	createdStudent, err := s.service.CreateStudent(ctx, student)
	if err != nil {
		logger.Error("creation failed", "error", err)
		return nil, s.convertError(err)
	}

	response := createdStudent.StudentToProto()
	logger.Info("student created", "student_id", createdStudent.ID.String(), "contract", createdStudent.ContractNumber)
	return response, nil
}
