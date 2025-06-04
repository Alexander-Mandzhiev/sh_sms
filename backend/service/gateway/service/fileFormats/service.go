package file_formats_service

import (
	library "backend/protos/gen/go/library"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

type FileFormatsService interface {
	FileFormatExists(ctx context.Context, req *library.FileFormatExistsRequest) (*library.FileFormatExistsResponse, error)
	ListFileFormats(ctx context.Context) (*library.ListFileFormatsResponse, error)
}

type fileFormatsService struct {
	client library.FileFormatServiceClient
	logger *slog.Logger
}

func NewFileFormatsService(client library.FileFormatServiceClient, logger *slog.Logger) FileFormatsService {
	return &fileFormatsService{
		client: client,
		logger: logger.With("service", "file format"),
	}
}

func (s *fileFormatsService) FileFormatExists(ctx context.Context, req *library.FileFormatExistsRequest) (*library.FileFormatExistsResponse, error) {
	s.logger.Debug("getting client", "format", req.Format)
	return s.client.FileFormatExists(ctx, req)
}

func (s *fileFormatsService) ListFileFormats(ctx context.Context) (*library.ListFileFormatsResponse, error) {
	s.logger.Debug("listing file formats")
	return s.client.ListFileFormats(ctx, &emptypb.Empty{})
}
