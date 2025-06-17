package teachers_service

import (
	"backend/pkg/models/teacher"
	"context"
	"github.com/google/uuid"
	"log/slog"
)

type TeachersProvider interface {
	CreateTeacher(ctx context.Context, teacher *teachers_models.CreateTeacher) (*teachers_models.Teacher, error)
	GetTeacher(ctx context.Context, id, clientID uuid.UUID) (*teachers_models.Teacher, error)
	UpdateTeacher(ctx context.Context, update *teachers_models.UpdateTeacher) (*teachers_models.Teacher, error)
	DeleteTeacher(ctx context.Context, id, clientID uuid.UUID) error
	ListTeachers(ctx context.Context, filter *teachers_models.ListTeachersFilter) (*teachers_models.ListTeachersResponse, error)
}

type Service struct {
	logger   *slog.Logger
	provider TeachersProvider
}

func New(provider TeachersProvider, logger *slog.Logger) *Service {
	const op = "service.New.PrivateSchool.Teachers"

	if logger == nil {
		logger = slog.Default()
	}

	logger.Info("initializing library handle - service teachers", slog.String("op", op))
	return &Service{
		provider: provider,
		logger:   logger,
	}
}
