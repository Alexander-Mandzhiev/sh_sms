package handle

import (
	"backend/protos/gen/go/apps"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) GenerateSecretKey(ctx context.Context, req *apps.GenerateSecretKeyRequest) (*apps.SecretKeyResponse, error) {
	if req.GetAppId() == "" || req.GetGeneratedBy() == "" {
		return nil, status.Error(codes.InvalidArgument, "app_id and generated_by are required")
	}
	return s.service.GenerateSecretKey(ctx, req)
}
