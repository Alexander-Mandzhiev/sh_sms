package handle

import (
	"backend/protos/gen/go/apps"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) Delete(ctx context.Context, req *apps.DeleteRequest) (*apps.DeleteResponse, error) {
	if req.GetAppId() == "" || req.GetDeletedBy() == "" {
		return nil, status.Error(codes.InvalidArgument, "app_id and deleted_by are required")
	}
	return s.service.Delete(ctx, req)
}
