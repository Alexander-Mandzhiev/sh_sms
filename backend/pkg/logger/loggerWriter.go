package sl

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"strings"
	"sync"
)

type LoggerWriter struct {
	logger *slog.Logger
	level  slog.Level
	buf    bytes.Buffer
	mu     sync.Mutex
}

func (w *LoggerWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	n, err = w.buf.Write(p)
	if err != nil {
		return n, err
	}

	for {
		line, err := w.buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				w.buf.WriteString(line)
				return n, nil
			}
			return n, err
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		switch w.level {
		case slog.LevelDebug:
			w.logger.Debug(line)
		case slog.LevelInfo:
			w.logger.Info(line)
		case slog.LevelWarn:
			w.logger.Warn(line)
		case slog.LevelError:
			w.logger.Error(line)
		default:
			w.logger.Log(context.Background(), w.level, line)
		}
	}
}

func NewLoggerWriter(logger *slog.Logger, level slog.Level) io.Writer {
	return &LoggerWriter{
		logger: logger,
		level:  level,
	}
}
