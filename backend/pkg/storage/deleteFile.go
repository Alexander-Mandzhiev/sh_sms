package storage

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (s *LocalStorage) DeleteFile(ctx context.Context, fileURL string) error {
	fullPath := filepath.Join(s.baseDir, fileURL)
	if !strings.HasPrefix(fullPath, s.baseDir) {
		return fmt.Errorf("invalid file path: %s", fileURL)
	}

	if err := os.Remove(fullPath); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("failed to delete file: %w", err)
	}
	return nil
}
