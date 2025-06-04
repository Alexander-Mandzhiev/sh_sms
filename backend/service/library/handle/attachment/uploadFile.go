package attachment_handle

import (
	"backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log/slog"
)

func (s *serverAPI) UploadFile(stream library.AttachmentService_UploadFileServer) error {
	const op = "grpc.Attachment.UploadFile"
	logger := s.logger.With(slog.String("op", op))
	logger.Info("Starting file upload")

	req, err := stream.Recv()
	if err != nil {
		logger.Error("Failed to receive metadata", "error", err)
		return status.Error(codes.InvalidArgument, "missing metadata")
	}

	meta := req.GetMetadata()
	if meta == nil {
		logger.Warn("First message must contain metadata")
		return status.Error(codes.InvalidArgument, "first message must contain metadata")
	}

	pr, pw := io.Pipe()
	defer pr.Close()

	saveErr := make(chan error, 1)
	var uploadedFile *library_models.UploadedFile

	go func() {
		defer close(saveErr)
		file, err := s.service.UploadFile(stream.Context(), &library_models.FileMetadata{
			BookID: meta.BookId,
			Format: meta.Format,
		}, pr)

		if err != nil {
			saveErr <- err
			return
		}

		uploadedFile = file
		saveErr <- nil
	}()

	var totalSize int64
	for {
		req, err = stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			pw.CloseWithError(err)
			return s.convertError(err)
		}

		chunk := req.GetChunkData()
		if chunk == nil {
			pw.CloseWithError(status.Error(codes.InvalidArgument, "expected chunk data"))
			return status.Error(codes.InvalidArgument, "expected chunk data")
		}

		if _, err = pw.Write(chunk); err != nil {
			logger.Error("Failed to write chunk", "error", err)
			return s.convertError(err)
		}
		totalSize += int64(len(chunk))
	}

	pw.Close()
	if err = <-saveErr; err != nil {
		logger.Error("File save failed", "error", err)
		return s.convertError(err)
	}

	logger.Info("File uploaded successfully", "book_id", meta.BookId, "format", meta.Format, "size", totalSize)

	return stream.SendAndClose(&library.UploadFileResponse{
		FileUrl:  uploadedFile.FilePath,
		Size:     uploadedFile.Size,
		Checksum: uploadedFile.Checksum,
	})
}
