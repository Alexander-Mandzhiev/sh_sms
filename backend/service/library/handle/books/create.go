package books_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"context"
	"log/slog"
)

func (s *serverAPI) CreateBook(ctx context.Context, req *library.CreateBookRequest) (*library.BookResponse, error) {
	const op = "grpc.Books.CreateBook"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", req.GetClientId()), slog.String("title", req.GetTitle()))
	logger.Debug("create book request received")

	params, err := library_models.CreateParamsFromProto(req)
	if err != nil {
		logger.Warn("invalid create parameters", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	bookResp, err := s.service.CreateBook(ctx, params)
	if err != nil {
		logger.Error("failed to create book", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	protoResponse := bookResp.BookResponseToProto()
	logger.Info("book created successfully", slog.Int64("id", protoResponse.GetId()), slog.String("title", protoResponse.GetTitle()))
	return protoResponse, nil
}
