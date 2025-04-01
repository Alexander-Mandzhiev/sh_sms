package handle

import (
	"backend/protos/gen/go/apps"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) Get(ctx context.Context, req *apps.GetRequest) (*apps.App, error) {
	if req.GetAppId() == "" {
		return nil, status.Error(codes.InvalidArgument, "app_id is required")
	}
	return s.service.Get(ctx, req)
}
