package sl

import (
	"io"
	"log/slog"
	"os"
)

const (
	EnvLocal      = "local"
	EnvDev        = "development"
	EnvTest       = "test"
	EnvStaging    = "staging"
	EnvProduction = "production"
)

type LoggerOptions struct {
	Output    io.Writer
	AddSource bool
}

func SetupLogger(env string, opts ...LoggerOptions) *slog.Logger {
	var opt LoggerOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	if opt.Output == nil {
		opt.Output = os.Stdout
	}

	level := getLogLevel(env)
	handlerOpts := &slog.HandlerOptions{
		Level:     level,
		AddSource: opt.AddSource && (env == EnvLocal || env == EnvDev),
	}

	var handler slog.Handler
	if env == EnvProduction || env == EnvStaging {
		handler = slog.NewJSONHandler(opt.Output, handlerOpts)
	} else {
		handler = slog.NewTextHandler(opt.Output, handlerOpts)
	}

	return slog.New(handler)
}

func getLogLevel(env string) slog.Level {
	switch env {
	case EnvLocal, EnvDev:
		return slog.LevelDebug
	case EnvTest:
		return slog.LevelWarn
	case EnvProduction, EnvStaging:
		return slog.LevelInfo
	default:
		return slog.LevelInfo
	}
}
