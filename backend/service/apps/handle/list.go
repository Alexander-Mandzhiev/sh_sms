package handle

import (
	"backend/protos/gen/go/apps"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) List(ctx context.Context, req *apps.ListRequest) (*apps.ListResponse, error) {
	if req.GetLimit() > 1000 {
		return nil, status.Error(codes.InvalidArgument, "maximum limit is 1000")
	}
	return s.service.List(ctx, req)
}
