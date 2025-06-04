package factory

import (
	library "backend/protos/gen/go/library"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

type LibraryClientType interface {
	library.AttachmentServiceClient
	library.BookServiceClient
	library.ClassServiceClient
	library.FileFormatServiceClient
	library.SubjectServiceClient
	Close() error
}

type LibraryClient struct {
	library.AttachmentServiceClient
	library.BookServiceClient
	library.ClassServiceClient
	library.FileFormatServiceClient
	library.SubjectServiceClient
	conn *grpc.ClientConn
}

func (c *LibraryClient) Close() error {
	return c.conn.Close()
}

func (p *ClientProvider) GetLibraryClient(ctx context.Context) (LibraryClientType, error) {
	client, err := p.getClient(ctx, ServiceLibrary)
	if err != nil {
		return nil, err
	}
	libraryClient, ok := client.(LibraryClientType)
	if !ok {
		return nil, fmt.Errorf("type assertion failed for SSO client")
	}
	return libraryClient, nil
}
