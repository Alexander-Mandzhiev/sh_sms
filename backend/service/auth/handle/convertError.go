package handle

import (
	"backend/protos/gen/go/auth"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func convertError(op string, err error) (*auth.AuthResponse, error) {
	switch {
	case errors.Is(err, ErrInvalidCredentials):
		return nil, status.Errorf(codes.Unauthenticated, "%s: %v", op, err)
	case errors.Is(err, ErrUserInactive):
		return nil, status.Errorf(codes.PermissionDenied, "%s: %v", op, err)
	case errors.Is(err, ErrClientAppInactive):
		return nil, status.Errorf(codes.PermissionDenied, "%s: %v", op, err)
	default:
		return nil, status.Errorf(codes.Internal, "%s: internal error", op)
	}
}
