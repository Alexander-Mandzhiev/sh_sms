package sso_models

import (
	"backend/protos/gen/go/sso/users_roles"
	"time"
)

type RevokeResponse struct {
	Success   bool      `json:"success"`
	RevokedAt time.Time `json:"revoked_at"`
}

func RevokeResponseFromProto(pb *user_roles.RevokeResponse) *RevokeResponse {
	return &RevokeResponse{
		Success:   pb.Success,
		RevokedAt: pb.RevokedAt.AsTime(),
	}
}
