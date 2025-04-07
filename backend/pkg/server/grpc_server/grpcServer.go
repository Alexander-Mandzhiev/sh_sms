package grpc_server

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"net"
)

type App struct {
	GRPCServer *grpc.Server
	listener   net.Listener
}

func New() *App {
	loggingOpts := []logging.Option{
		logging.WithLogOnEvents(
			logging.PayloadReceived, logging.PayloadSent,
		),
	}
	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			slog.Error("Recovered from panic", slog.Any("panic", p))
			return status.Errorf(codes.Internal, "internal error")
		}),
	}
	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(recoveryOpts...),
		logging.UnaryServerInterceptor(InterceptorLogger(slog.Default()), loggingOpts...),
	))

	return &App{
		GRPCServer: gRPCServer,
	}
}

func InterceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

func (a *App) MustRun(logger *slog.Logger, port int) {
	if err := a.Run(port); err != nil {
		logger.Error("Failed to start GRPC server", slog.Any("error", err))
	}
}

func (a *App) Run(port int) error {
	const op = "grpcapp.Run"
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	a.listener = listener

	slog.Info("gRPC server started", slog.String("addr", listener.Addr().String()))

	if err = a.GRPCServer.Serve(listener); err != nil {
		if err == grpc.ErrServerStopped || err.Error() == "use of closed network connection" {
			slog.Info("gRPC server stopped gracefully")
			return nil
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Shutdown() {
	slog.Info("Stopping gRPC server")
	if a.listener != nil {
		if err := a.listener.Close(); err != nil {
			slog.Error("Failed to close listener", slog.Any("error", err))
		}
	}
	a.GRPCServer.GracefulStop()
}
