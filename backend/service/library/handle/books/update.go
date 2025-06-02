package books_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"context"
	"log/slog"
)

func (s *serverAPI) UpdateBook(ctx context.Context, req *library.UpdateBookRequest) (*library.BookResponse, error) {
	const op = "grpc.Library.Books.UpdateBook"
	logger := s.logger.With(slog.String("op", op), slog.Int64("id", req.GetId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("Update book called")

	updateReq, err := library_models.UpdateBookRequestFromProto(req)
	if err != nil {
		logger.Warn("Invalid update parameters", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	bookResp, err := s.service.UpdateBook(ctx, updateReq)
	if err != nil {
		logger.Error("Failed to update book", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	protoResponse := bookResp.BookResponseToProto()
	logger.Info("Book updated", slog.Int64("id", protoResponse.GetId()))
	return protoResponse, nil
}
