package handle

import (
	"backend/protos/gen/go/apps"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) RotateSecretKey(ctx context.Context, req *apps.RotateSecretKeyRequest) (*apps.SecretKeyResponse, error) {
	if req.GetAppId() == "" || req.GetRotatedBy() == "" {
		return nil, status.Error(codes.InvalidArgument, "app_id and rotated_by are required")
	}
	return s.service.RotateSecretKey(ctx, req)
}
