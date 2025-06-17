package students_handle

import (
	"backend/pkg/models/students"
	"backend/protos/gen/go/private_school"
	"context"
	"log/slog"
)

func (s *serverAPI) ListStudents(ctx context.Context, req *private_school_v1.ListStudentsRequest) (*private_school_v1.ListStudentsResponse, error) {
	const op = "grpc.PrivateSchool.StudentService.ListStudents"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", req.GetClientId()))
	logger.Debug("ListStudents called")

	params, err := students_models.ListStudentsParamsFromProto(req)
	if err != nil {
		logger.Warn("Invalid list parameters", slog.String("error", err.Error()))
		return nil, s.convertError(err)
	}

	studentsList, nextCursor, err := s.service.ListStudents(ctx, params)
	if err != nil {
		logger.Error("Failed to list students", slog.String("error", err.Error()))
		return nil, s.convertError(err)
	}

	domainResponse := &students_models.ListStudentsResponse{
		Students:   studentsList,
		NextCursor: nextCursor,
	}

	logger.Info("Students listed successfully", slog.Int("count", len(studentsList)), slog.String("next_cursor", nextCursor))
	return domainResponse.ToProto(), nil
}
