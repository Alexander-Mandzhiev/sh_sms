package books_service

import (
	"backend/pkg/models/library"
	"context"
	"errors"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) GetBook(ctx context.Context, id int64, clientID uuid.UUID) (*library_models.BookResponse, error) {
	const op = "service.Library.Books.GetBook"
	logger := s.logger.With(slog.String("op", op), slog.Int64("id", id), slog.String("client_id", clientID.String()))
	logger.Debug("Getting book")

	book, err := s.provider.GetBookByID(ctx, id, clientID)
	if err != nil {
		logger.Error("Failed to get book", slog.Any("error", err))
		return nil, err
	}

	subject, err := s.subjectsProvide.GetSubjectByID(ctx, int32(book.SubjectID))
	if err != nil {
		if errors.Is(err, library_models.ErrNotFound) {
			logger.Warn("Subject does not exist", slog.Int("subject_id", book.SubjectID))
			return nil, library_models.ErrBookInvalidSubjectID
		}
		logger.Error("Failed to check subject existence", slog.Int("subject_id", book.SubjectID), slog.Any("error", err))
		return nil, err
	}

	class, err := s.classesProvider.GetClassByID(ctx, book.ClassID)
	if err != nil {
		if errors.Is(err, library_models.ErrNotFound) {
			logger.Warn("class does not exist", slog.Int("class_id", book.ClassID), slog.Any("error", err))
			return nil, library_models.ErrBookInvalidClassID
		}
		logger.Error("failed to get class", slog.Int("class_id", book.ClassID), slog.Any("error", err))
		return nil, err
	}

	logger.Debug("Book retrieved", slog.String("title", book.Title))
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
