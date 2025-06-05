package storage

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func (s *LocalStorage) GetFile(ctx context.Context, fileID string) (io.ReadSeekCloser, string, int64, error) {
	fullPath := filepath.Join(s.baseDir, fileID)
	file, err := os.Open(fullPath)
	if err != nil {
		return nil, "", 0, fmt.Errorf("failed to open file: %w", err)
	}

	stat, err := file.Stat()
	if err != nil {
		file.Close()
		return nil, "", 0, fmt.Errorf("failed to get file info: %w", err)
	}

	return file, fullPath, stat.Size(), nil
}
