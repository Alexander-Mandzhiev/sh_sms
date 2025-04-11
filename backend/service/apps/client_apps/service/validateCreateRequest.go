package service

import (
	pb "backend/protos/gen/go/apps/clients_apps"
	"fmt"
	"github.com/google/uuid"
)

func validateCreateRequest(req *pb.CreateRequest) error {
	if _, err := uuid.Parse(req.ClientId); err != nil {
		return fmt.Errorf("%w: invalid client_id format", ErrInvalidArgument)
	}

	if req.AppId <= 0 {
		return fmt.Errorf("%w: app_id must be positive", ErrInvalidArgument)
	}
	return nil
}
