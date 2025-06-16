package teachers_service

import (
	"backend/pkg/models/private_school"
	"context"
	"github.com/google/uuid"
	"log/slog"
)

type TeachersProvider interface {
	CreateTeacher(ctx context.Context, teacher *private_school_models.CreateTeacher) (*private_school_models.Teacher, error)
	GetTeacher(ctx context.Context, id, clientID uuid.UUID) (*private_school_models.Teacher, error)
	UpdateTeacher(ctx context.Context, update *private_school_models.UpdateTeacher) (*private_school_models.Teacher, error)
	DeleteTeacher(ctx context.Context, id, clientID uuid.UUID) error
	ListTeachers(ctx context.Context, filter *private_school_models.ListTeachersFilter) (*private_school_models.ListTeachersResponse, error)
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
