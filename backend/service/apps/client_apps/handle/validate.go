package handle

import (
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validateClientID(clientID string) error {
	if clientID == "" {
		return status.Error(codes.InvalidArgument, "client_id is required")
	}
	if _, err := uuid.Parse(clientID); err != nil {
		return status.Error(codes.InvalidArgument, "invalid client_id format")
	}
	return nil
}

func validateAppID(appID int32) error {
	if appID <= 0 {
		return status.Error(codes.InvalidArgument, "app_id must be positive")
	}
	return nil
}

func validatePagination(page, count int32) error {
	if page <= 0 || count <= 0 {
		return status.Error(codes.InvalidArgument, "page and count must be positive")
	}
	return nil
}
