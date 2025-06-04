package attachments_service

import (
	library "backend/protos/gen/go/library"
	"context"
	"io"
)

func (s *attachmentsService) UploadFile(ctx context.Context, metadata *library.FileMetadata, file io.Reader) (*library.UploadFileResponse, error) {
	s.logger.Debug("Uploading file", "book_id", metadata.BookId, "format", metadata.Format)

	stream, err := s.client.UploadFile(ctx)
	if err != nil {
		return nil, err
	}

	if err = stream.Send(&library.UploadFileRequest{Request: &library.UploadFileRequest_Metadata{Metadata: metadata}}); err != nil {
		return nil, err
	}

	buf := make([]byte, 1024*32)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if err := stream.Send(&library.UploadFileRequest{Request: &library.UploadFileRequest_ChunkData{ChunkData: buf[:n]}}); err != nil {
			return nil, err
		}
	}

	return stream.CloseAndRecv()
}
