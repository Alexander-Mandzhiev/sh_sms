package sl

import (
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
)

func LogGinRequest(logger *slog.Logger, params gin.LogFormatterParams, level slog.Level) {
	if logger == nil {
		return
	}

	attrs := []slog.Attr{
		slog.String("method", params.Method),
		slog.String("path", params.Path),
		slog.String("full_path", params.Request.URL.String()),
		slog.Int("status", params.StatusCode),
		slog.Duration("latency", params.Latency),
		slog.String("client_ip", params.ClientIP),
		slog.String("user_agent", params.Request.UserAgent()),
		slog.Time("time", params.TimeStamp),
	}

	if params.ErrorMessage != "" {
		attrs = append(attrs, slog.String("error", params.ErrorMessage))
	}

	if requestID := params.Request.Header.Get("X-Request-ID"); requestID != "" {
		attrs = append(attrs, slog.String("request_id", requestID))
	}

	logger.LogAttrs(
		context.Background(),
		level,
		"HTTP request",
		attrs...,
	)
}
