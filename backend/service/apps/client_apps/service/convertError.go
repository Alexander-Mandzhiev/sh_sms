package service

import (
	"backend/service/apps/constants"
	"errors"
)

func (s *Service) convertError(err error) error {
	switch {
	case errors.Is(err, constants.ErrAlreadyExists):
		return constants.ErrAlreadyExists
	case errors.Is(err, constants.ErrNotFound):
		return constants.ErrNotFound
	case errors.Is(err, constants.ErrPermissionDenied):
		return constants.ErrPermissionDenied
	default:
		return constants.ErrInternal
	}
}
