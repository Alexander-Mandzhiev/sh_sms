package storage

import (
	"backend/pkg/models/library"
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func (s *LocalStorage) SaveFile(ctx context.Context, meta library_models.FileMetadata, r io.Reader) (library_models.UploadedFile, error) {
	fileName := fmt.Sprintf("book_%d_%s.%s", meta.BookID, generateUUID(), meta.Format)
	filePath := filepath.Join(s.baseDir, fileName)

	f, err := os.Create(filePath)
	if err != nil {
		return library_models.UploadedFile{}, fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	hasher := sha256.New()
	limitedReader := io.LimitReader(r, s.maxFileSize)
	multiWriter := io.MultiWriter(f, hasher)

	size, err := io.Copy(multiWriter, limitedReader)
	if err != nil {
		os.Remove(filePath)
		return library_models.UploadedFile{}, fmt.Errorf("failed to write file: %w", err)
	}

	if size >= s.maxFileSize {
		os.Remove(filePath)
		return library_models.UploadedFile{}, fmt.Errorf("file size exceeds limit of %d MB", s.maxFileSize/(1024*1024))
	}

	return library_models.UploadedFile{
		FilePath: fileName,
		Size:     size,
		Checksum: fmt.Sprintf("%x", hasher.Sum(nil)),
	}, nil
}
