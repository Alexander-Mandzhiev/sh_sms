package service

import (
	"backend/service/apps/client_apps/handle"
	"github.com/google/uuid"
)

func validateClientID(clientID string) error {
	if clientID == "" {
		return handle.ErrInvalidArgument
	}
	if _, err := uuid.Parse(clientID); err != nil {
		return handle.ErrInvalidArgument
	}
	return nil
}

func validateAppID(appID int) error {
	if appID <= 0 {
		return handle.ErrInvalidArgument
	}
	return nil
}

func validatePagination(page, count int) error {
	if page <= 0 || count <= 0 {
		return handle.ErrInvalidArgument
	}
	if count > 1000 {
		return handle.ErrInvalidArgument
	}
	return nil
}
