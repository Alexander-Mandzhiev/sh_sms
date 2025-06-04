package books_service

import (
	"backend/pkg/models/library"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) UpdateBook(ctx context.Context, req *library_models.UpdateBookRequest) (*library_models.BookResponse, error) {
	const op = "service.Library.Books.Update"
	logger := s.logger.With(slog.String("op", op), slog.Int64("book_id", req.ID), slog.String("client_id", req.ClientID.String()))
	logger.Debug("update book request received")

	book, err := s.provider.GetBookByID(ctx, req.ID, req.ClientID)
	if err != nil {
		if errors.Is(err, library_models.ErrNotFound) {
			logger.Warn("book not found", slog.Any("error", err))
			return nil, library_models.ErrNotFound
		}
		logger.Error("failed to get book", slog.Any("error", err))
		return nil, err
	}
	logger.Debug("book retrieved from provider", slog.Any("current_book", book))

	applyUpdates(book, req)
	logger.Debug("updates applied to book model")

	updatedBook, err := s.provider.UpdateBook(ctx, book)
	if err != nil {
		logger.Error("failed to update book in provider", slog.Any("error", err))
		return nil, err
	}
	logger.Debug("book updated in provider", slog.Any("updated_book", updatedBook))

	subject, err := s.subjectsProvide.GetSubjectByID(ctx, int32(updatedBook.SubjectID))
	if err != nil {
		if errors.Is(err, library_models.ErrNotFound) {
			logger.Warn("subject does not exist", slog.Int("subject_id", updatedBook.SubjectID), slog.Any("error", err))
			return nil, library_models.ErrBookInvalidSubjectID
		}
		logger.Error("failed to get subject", slog.Int("subject_id", updatedBook.SubjectID), slog.Any("error", err))
		return nil, err
	}
	logger.Debug("subject retrieved", slog.String("subject_name", subject.Name))

	class, err := s.classesProvider.GetClassByID(ctx, updatedBook.ClassID)
	if err != nil {
		if errors.Is(err, library_models.ErrNotFound) {
			logger.Warn("class does not exist", slog.Int("class_id", updatedBook.ClassID), slog.Any("error", err))
			return nil, library_models.ErrBookInvalidClassID
		}
		logger.Error("failed to get class", slog.Int("class_id", updatedBook.ClassID), slog.Any("error", err))
		return nil, err
	}
	logger.Debug("class retrieved", slog.Int("grade", int(class.Grade)))

	logger.Info("book updated successfully", slog.Int64("book_id", updatedBook.ID), slog.String("title", updatedBook.Title))
	return &library_models.BookResponse{
		ID:          updatedBook.ID,
		ClientID:    updatedBook.ClientID.String(),
		Title:       updatedBook.Title,
		Author:      updatedBook.Author,
		Description: updatedBook.Description,
		SubjectName: subject.Name,
		Grade:       class.Grade,
		CreatedAt:   updatedBook.CreatedAt,
	}, nil
}

func applyUpdates(book *library_models.Book, req *library_models.UpdateBookRequest) {
	if req.Title != nil {
		book.Title = *req.Title
	}
	if req.Author != nil {
		book.Author = *req.Author
	}
	if req.Description != nil {
		book.Description = *req.Description
	}
	if req.SubjectID != nil {
		book.SubjectID = int(*req.SubjectID)
	}
	if req.ClassID != nil {
		book.ClassID = int(*req.ClassID)
	}
}
