package storage

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func (s *LocalStorage) GetFile(ctx context.Context, fileURL string) (io.ReadSeeker, string, int64, error) {
	cleanPath := filepath.Clean(fileURL)
	fullPath := filepath.Join(s.baseDir, cleanPath)

	if !strings.HasPrefix(fullPath, s.baseDir) {
		return nil, "", 0, fmt.Errorf("invalid file path")
	}

	file, err := os.Open(fullPath)
	if err != nil {
		return nil, "", 0, err
	}

	stat, err := file.Stat()
	if err != nil {
		file.Close()
		return nil, "", 0, err
	}

	return file, fullPath, stat.Size(), nil
}
