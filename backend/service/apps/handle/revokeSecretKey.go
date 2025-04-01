package handle

import (
	"backend/protos/gen/go/apps"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) RevokeSecretKey(ctx context.Context, req *apps.RevokeSecretKeyRequest) (*apps.RevokeSecretKeyResponse, error) {
	if req.GetAppId() == "" || req.GetRevokedBy() == "" {
		return nil, status.Error(codes.InvalidArgument, "app_id and revoked_by are required")
	}
	return s.service.RevokeSecretKey(ctx, req)
}
