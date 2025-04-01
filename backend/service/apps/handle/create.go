package handle

import (
	"backend/protos/gen/go/apps"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) Create(ctx context.Context, req *apps.CreateRequest) (*apps.App, error) {
	if req.GetName() == "" || req.GetCreatedBy() == "" {
		return nil, status.Error(codes.InvalidArgument, "name and created_by are required")
	}
	return s.service.Create(ctx, req)
}
