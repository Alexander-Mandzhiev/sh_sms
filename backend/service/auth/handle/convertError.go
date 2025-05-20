package handle

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

var (
	ErrInvalidToken        = errors.New("invalid token")
	ErrSessionNotFound     = errors.New("session not found")
	ErrTokenAlreadyRevoked = errors.New("token already revoked")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrUserInactive        = errors.New("user is inactive")
	ErrClientAppInactive   = errors.New("client app is inactive")
	ErrSessionExpired      = errors.New("session expired")
	ErrInvalidClientOrApp  = errors.New("invalid client or app combination")
	ErrNoRowsAffected      = errors.New("no rows affected by operation")
)

func (h *serverAPI) convertError(op string, err error) error {
	switch {
	case errors.Is(err, ErrInvalidCredentials),
		errors.Is(err, ErrSessionExpired),
		errors.Is(err, ErrInvalidToken):
		return status.Errorf(codes.Unauthenticated, "%s: %v", op, err)

	case errors.Is(err, ErrUserInactive),
		errors.Is(err, ErrClientAppInactive),
		errors.Is(err, ErrInvalidClientOrApp):
		return status.Errorf(codes.PermissionDenied, "%s: %v", op, err)

	case errors.Is(err, ErrSessionNotFound):
		return status.Errorf(codes.NotFound, "%s: %v", op, err)

	case errors.Is(err, ErrTokenAlreadyRevoked),
		errors.Is(err, ErrNoRowsAffected):
		return status.Errorf(codes.FailedPrecondition, "%s: %v", op, err)

	default:
		h.logger.Error("internal error", slog.String("op", op), slog.Any("error", err))
		return status.Errorf(codes.Internal, "%s: internal error", op)
	}
}
