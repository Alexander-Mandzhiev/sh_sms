package storage

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (s *LocalStorage) DeleteFile(ctx context.Context, fileURL string) error {
	cleanPath := filepath.Clean(fileURL)
	fullPath := filepath.Join(s.baseDir, cleanPath)
	relPath, err := filepath.Rel(s.baseDir, fullPath)
	if err != nil {
		return fmt.Errorf("invalid file path: %w", err)
	}

	if strings.HasPrefix(relPath, "..") {
		return fmt.Errorf("invalid file path: %s", fileURL)
	}

	if err = os.Remove(fullPath); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("failed to delete file: %w", err)
	}
	return nil
}
