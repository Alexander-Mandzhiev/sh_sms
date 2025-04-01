package handle

import (
	"backend/protos/gen/go/apps"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) GetKeyRotationHistory(ctx context.Context, req *apps.GetKeyRotationHistoryRequest) (*apps.KeyRotationHistoryResponse, error) {
	if req.GetAppId() == "" {
		return nil, status.Error(codes.InvalidArgument, "app_id is required")
	}
	if req.GetLimit() != nil && req.GetLimit().GetValue() > 1000 {
		return nil, status.Error(codes.InvalidArgument, "maximum limit is 1000")
	}
	return s.service.GetKeyRotationHistory(ctx, req)
}
