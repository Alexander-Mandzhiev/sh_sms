package library_models

import library "backend/protos/gen/go/library"

type FileMetadata struct {
	BookID int64  `json:"book_id"`
	Format string `json:"format"`
}

func (m *FileMetadata) ToProto() *library.FileMetadata {
	return &library.FileMetadata{
		BookId: m.BookID,
		Format: m.Format,
	}
}

func FileMetadataFromProto(proto *library.FileMetadata) *FileMetadata {
	return &FileMetadata{
		BookID: proto.BookId,
		Format: proto.Format,
	}
}
