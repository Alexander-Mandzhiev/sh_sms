package students_handle

import (
	"backend/pkg/models/students"
	"backend/pkg/utils"
	"backend/protos/gen/go/private_school"
	"context"
	"log/slog"
)

func (s *serverAPI) ListStudents(ctx context.Context, req *private_school_v1.ListStudentsRequest) (*private_school_v1.ListStudentsResponse, error) {
	const op = "grpc.StudentService.ListStudents"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", req.GetClientId()), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("list students request")

	params, err := students_models.ListStudentsRequestFromProto(req)
	if err != nil {
		logger.Warn("invalid list parameters", "error", err)
		return nil, s.convertError(err)
	}

	studentsList, nextCursor, err := s.service.ListStudents(ctx, params)
	if err != nil {
		logger.Error("failed to list students", "error", err)
		return nil, s.convertError(err)
	}

	response := &private_school_v1.ListStudentsResponse{
		Students: make([]*private_school_v1.StudentResponse, len(studentsList)),
	}

	for i, student := range studentsList {
		response.Students[i] = student.StudentToProto()
	}

	if nextCursor != nil {
		response.NextCursor = nextCursor.ToProto()
	}

	logger.Info("students listed", "count", len(studentsList))
	return response, nil
}
