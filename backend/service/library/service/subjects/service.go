package subjects_service

import (
	"backend/pkg/models/library"
	"context"
	"log/slog"
)

type SubjectsProvider interface {
	CreateSubject(ctx context.Context, subject *library_models.Subject) (int32, error)
	GetSubjectByID(ctx context.Context, id int32) (*library_models.Subject, error)
	UpdateSubject(ctx context.Context, subject *library_models.Subject) error
	DeleteSubject(ctx context.Context, id int32) error
	ListSubjects(ctx context.Context) ([]*library_models.Subject, error)
}
type Service struct {
	logger   *slog.Logger
	provider SubjectsProvider
}

func New(provider SubjectsProvider, logger *slog.Logger) *Service {
	const op = "service.New.Library.Subjects"

	if logger == nil {
		logger = slog.Default()
	}

	logger.Info("initializing library handle - service subjects", slog.String("op", op))
	return &Service{
		provider: provider,
		logger:   logger,
	}
}
