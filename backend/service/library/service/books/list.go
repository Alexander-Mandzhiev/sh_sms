package books_service

import (
	"backend/pkg/models/library"
	"context"
	"log/slog"
)

func (s *Service) ListBooks(ctx context.Context, params *library_models.ListBooksRequest) (*library_models.ListBooksResponse, error) {
	const op = "service.Library.Books.List"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", params.ClientID.String()), slog.Int("count", int(params.Count)))
	if params.Cursor != nil {
		logger = logger.With(slog.Int64("cursor", *params.Cursor))
	}
	logger.Debug("processing list books request")

	books, nextCursor, err := s.provider.ListBooks(ctx, params)
	if err != nil {
		logger.Error("failed to get books list", slog.Any("error", err))
		return nil, err
	}

	if len(books) == 0 {
		logger.Debug("no books found")
		return &library_models.ListBooksResponse{
			Books:      []*library_models.BookResponse{},
			NextCursor: 0,
			TotalCount: 0,
		}, nil
	}

	total, err := s.provider.CountBooks(ctx, params.ClientID, params.Filter)
	if err != nil {
		logger.Error("failed to count books", slog.Any("error", err))
		return nil, err
	}

	subjects, err := s.subjectsProvide.ListSubjects(ctx)
	if err != nil {
		logger.Error("Failed to getting list subjects", slog.String("client_id", params.ClientID.String()), slog.Any("error", err))
		return nil, err
	}

	classes, err := s.classesProvider.ListClasses(ctx)
	if err != nil {
		logger.Error("failed to getting list class", slog.String("client_id", params.ClientID.String()), slog.Any("error", err))
		return nil, err
	}

	subjectMap := make(map[int32]string)
	for _, sub := range subjects {
		subjectMap[sub.ID] = sub.Name
	}

	classMap := make(map[int32]int32)
	for _, c := range classes {
		classMap[c.ID] = c.Grade
	}

	var enrichedBooks []*library_models.BookResponse
	for _, book := range books {
		enrichBook := &library_models.BookResponse{
			ID:          book.ID,
			ClientID:    book.ClientID.String(),
			Title:       book.Title,
			Author:      book.Author,
			Description: book.Description,
			SubjectName: subjectMap[int32(book.SubjectID)],
			Grade:       classMap[int32(book.ClassID)],
			CreatedAt:   book.CreatedAt,
		}
		enrichedBooks = append(enrichedBooks, enrichBook)

	}

	logger.Info("Books listing completed", slog.Int("books_count", len(books)), slog.Int64("has_next_page", nextCursor))
	return &library_models.ListBooksResponse{
		Books:      enrichedBooks,
		NextCursor: nextCursor,
		TotalCount: total,
	}, nil
}
