// pkg/storage/file_storage.go
package storage

import (
	"backend/pkg/config/models"
	"backend/pkg/models/library"
	"context"
	"fmt"
	"github.com/google/uuid"
	"io"
	"os"
)

type FileStorage interface {
	SaveFile(ctx context.Context, meta library_models.FileMetadata, r io.Reader) (library_models.UploadedFile, error)
	DeleteFile(ctx context.Context, fileURL string) error
}

type LocalStorage struct {
	baseDir     string
	maxFileSize int64
	baseURL     string
}

func NewLocalStorage(cfg models.FileStorageConfig) *LocalStorage {
	if err := os.MkdirAll(cfg.BaseDir, 0755); err != nil {
		panic(fmt.Sprintf("failed to create storage directory: %v", err))
	}

	return &LocalStorage{
		baseDir:     cfg.BaseDir,
		maxFileSize: int64(cfg.MaxFileSizeMB) * 1024 * 1024,
		baseURL:     cfg.BaseURL,
	}
}

func generateUUID() string {
	return uuid.New().String()
}
