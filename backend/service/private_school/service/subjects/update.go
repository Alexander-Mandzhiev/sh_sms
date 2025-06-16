package subjects_service

import (
	private_school_models "backend/pkg/models/private_school"
	"context"
	"log/slog"
)

func (s *Service) UpdateSubject(ctx context.Context, params *private_school_models.UpdateSubjectParams) (*private_school_models.Subject, error) {
	const op = "service.PrivateSchool.Subjects.UpdateSubject"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("updating subject", slog.Int("id", int(params.ID)), slog.String("name", params.Name))

	if params.ID <= 0 {
		logger.Warn("invalid subject ID")
		return nil, private_school_models.ErrInvalidSubjectID
	}

	currentSubject, err := s.GetSubject(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	if currentSubject.Name != params.Name {
		currentSubject.Name = params.Name
		if err = s.provider.UpdateSubject(ctx, currentSubject); err != nil {
			logger.Error("failed to update subject", slog.String("error", err.Error()), slog.Int("id", int(params.ID)))
			return nil, err
		}
	} else {
		logger.Debug("subject data not changed, skip update")
	}

	logger.Info("subject updated successfully", slog.Int("id", int(params.ID)), slog.String("name", params.Name))
	return currentSubject, nil
}
