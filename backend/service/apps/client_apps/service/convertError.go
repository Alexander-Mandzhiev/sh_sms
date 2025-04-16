package service

import (
	"backend/service/apps/client_apps/handle"
	"errors"
)

func (s *Service) convertError(err error) error {
	switch {
	case errors.Is(err, handle.ErrAlreadyExists):
		return handle.ErrAlreadyExists
	case errors.Is(err, handle.ErrNotFound):
		return handle.ErrNotFound
	case errors.Is(err, handle.ErrPermissionDenied):
		return handle.ErrPermissionDenied
	default:
		return handle.ErrInternal
	}
}
