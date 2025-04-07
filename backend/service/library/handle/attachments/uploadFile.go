package attachments_handle

import (
	"backend/protos/gen/go/library/attachments"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

func (s *serverAPI) UploadFile(stream attachments.AttachmentsService_UploadFileServer) error {
	var metadata *attachments.FileMetadata
	chunks := make([][]byte, 0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Error(codes.Internal, "upload failed")
		}

		switch data := req.Data.(type) {
		case *attachments.UploadFileRequest_Metadata:
			metadata = data.Metadata
		case *attachments.UploadFileRequest_Chunk:
			chunks = append(chunks, data.Chunk)
		}
	}

	s.logger.Info("File upload completed", "name", metadata.FileName)
	return stream.SendAndClose(&attachments.Attachment{
		FileName: metadata.FileName,
	})
}
