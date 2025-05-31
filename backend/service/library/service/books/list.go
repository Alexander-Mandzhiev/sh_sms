package books_service

import (
	library_models "backend/pkg/models/library"
	"context"
	"fmt"
	"log/slog"
	"sync"
)

func (s *Service) ListBooks(ctx context.Context, params *library_models.ListBooksParams) (*library_models.ListBooksResult, error) {
	const op = "service.Library.Books.ListBooks"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", params.ClientID), slog.Int("page_size", int(params.PageSize)))
	logger.Debug("Processing list books request")

	if params.ClientID == "" {
		return nil, library_models.ErrClientIDRequired
	}
	if params.PageSize < 0 {
		return nil, fmt.Errorf("%w: %d", library_models.ErrInvalidPageSize, params.PageSize)
	}

	const defaultPageSize = 50
	const maxPageSize = 100
	if params.PageSize == 0 {
		params.PageSize = defaultPageSize
	} else if params.PageSize > maxPageSize {
		params.PageSize = maxPageSize
	}

	var books []*library_models.Book
	var total int32
	var wg sync.WaitGroup
	var booksErr, countErr error

	wg.Add(2)

	go func() {
		defer wg.Done()
		books, booksErr = s.provider.ListBooks(ctx, params)
	}()

	go func() {
		defer wg.Done()
		total, countErr = s.provider.CountBooks(ctx, params.ClientID, params.Filter)
	}()

	wg.Wait()

	if booksErr != nil {
		logger.Error("Failed to get books list", slog.Any("error", booksErr))
		return nil, booksErr
	}
	if countErr != nil {
		logger.Error("Failed to count books", slog.Any("error", countErr))
		return nil, countErr
	}

	var nextPageToken string
	if len(books) > 0 && len(books) == int(params.PageSize) {
		lastBook := books[len(books)-1]
		cursor := library_models.NewCursor(lastBook.ID, lastBook.CreatedAt)
		nextPageToken, _ = cursor.Encode()
	}

	result := &library_models.ListBooksResult{
		Books:         books,
		NextPageToken: nextPageToken,
		TotalCount:    total,
	}

	logger.Info("Books listing completed", slog.Int("books_count", len(books)), slog.Bool("has_next_page", nextPageToken != ""))
	return result, nil
}
