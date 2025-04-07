package attachments_handle

import (
	"backend/protos/gen/go/library/attachments"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

func (s *serverAPI) DownloadFile(req *attachments.GetAttachmentRequest, stream attachments.AttachmentsService_DownloadFileServer) error {
	s.logger.Info("Download request received", "id", req.Id)

	if err := stream.Send(&attachments.DownloadFileResponse{
		Data: &attachments.DownloadFileResponse_Metadata{
			Metadata: &attachments.FileMetadata{
				FileName: "example.txt",
			},
		},
	}); err != nil {
		return err
	}

	chunk := make([]byte, 1024)
	for {
		n, err := io.ReadFull(stream.Context().(io.Reader), chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Error(codes.Internal, "download failed")
		}
		stream.Send(&attachments.DownloadFileResponse{
			Data: &attachments.DownloadFileResponse_Chunk{
				Chunk: chunk[:n],
			},
		})
	}

	return nil
}
