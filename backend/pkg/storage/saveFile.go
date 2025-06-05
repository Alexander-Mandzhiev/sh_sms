package storage

import (
	"backend/pkg/models/library"
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func (s *LocalStorage) SaveFile(ctx context.Context, meta library_models.FileMetadata, r io.Reader) (library_models.UploadedFile, error) {
	safeFileName := sanitizeFileName(fmt.Sprintf("book_%s.%s", generateUUID(), meta.Format))
	filePath := filepath.Join(s.baseDir, safeFileName)

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
		FilePath: safeFileName,
		Size:     size,
		Checksum: fmt.Sprintf("%x", hasher.Sum(nil)),
	}, nil
}

func sanitizeFileName(name string) string {
	return strings.Map(func(r rune) rune {
		if r < 32 || r == '\\' || r == '/' || r == ':' || r == '*' ||
			r == '?' || r == '"' || r == '<' || r == '>' || r == '|' {
			return -1
		}
		return r
	}, name)
}
