package books_service

import (
	sl "backend/pkg/logger"
	"backend/pkg/models/library"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) CreateBook(ctx context.Context, params *library_models.CreateBookParams) (*library_models.BookResponse, error) {
	const op = "books_service.CreateBook"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", params.ClientID.String()), slog.String("title", params.Title))
	logger.Debug("creating book")

	logger.Debug("book parameters", slog.Int("subject_id", int(params.SubjectID)), slog.Int("class_id", int(params.ClassID)), slog.Int("title_len", len(params.Title)), slog.Int("author_len", len(params.Author)), slog.Int("desc_len", len(params.Description)))

	subject, err := s.subjectsProvide.GetSubjectByID(ctx, params.SubjectID)
	if err != nil {
		if errors.Is(err, library_models.ErrNotFound) {
			logger.Warn("subject not found", slog.Int("subject_id", int(params.SubjectID)))
			return nil, library_models.ErrBookInvalidSubjectID
		}
		logger.Error("failed to get subject", slog.Int("subject_id", int(params.SubjectID)), sl.Err(err, true))
		return nil, err
	}
	logger.Debug("subject found", slog.String("subject_name", subject.Name))

	class, err := s.classesProvider.GetClassByID(ctx, int(params.ClassID))
	if err != nil {
		if errors.Is(err, library_models.ErrNotFound) {
			logger.Warn("class not found", slog.Int("class_id", int(params.ClassID)))
			return nil, library_models.ErrBookInvalidClassID
		}
		logger.Error("failed to get class", slog.Int("class_id", int(params.ClassID)), sl.Err(err, true))
		return nil, err
	}
	logger.Debug("class found", slog.Int("grade", int(class.Grade)))

	book := &library_models.Book{
		ClientID:    params.ClientID,
		Title:       params.Title,
		Author:      params.Author,
		Description: params.Description,
		SubjectID:   int(subject.ID),
		ClassID:     int(class.ID),
	}

	if err = s.provider.CreateBook(ctx, book); err != nil {
		logger.Error("failed to create book in provider", sl.Err(err, true))
		return nil, err
	}
	logger.Debug("book created in provider", slog.Int64("id", book.ID))

	logger.Info("book created successfully", slog.Int64("book_id", book.ID), slog.String("title", book.Title))
	return &library_models.BookResponse{
		ID:          book.ID,
		ClientID:    book.ClientID.String(),
		Title:       book.Title,
		Author:      book.Author,
		Description: book.Description,
		SubjectName: subject.Name,
		Grade:       class.Grade,
		CreatedAt:   book.CreatedAt,
	}, nil
}
