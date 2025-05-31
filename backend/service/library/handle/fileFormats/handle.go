package file_format_handle

import (
	library "backend/protos/gen/go/library"
	"context"
	"google.golang.org/grpc"
	"log/slog"
)

type FileFormatService interface {
	ListFileFormats(ctx context.Context) ([]string, error)
	FileFormatExists(ctx context.Context, format string) (bool, error)
}

type serverAPI struct {
	library.UnimplementedFileFormatServiceServer
	service FileFormatService
	logger  *slog.Logger
}

func Register(gRPCServer *grpc.Server, service FileFormatService, logger *slog.Logger) {
	library.RegisterFileFormatServiceServer(gRPCServer, &serverAPI{
		service: service,
		logger:  logger.With("module", "file_formats"),
	})
}
