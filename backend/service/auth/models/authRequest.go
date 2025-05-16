package models

import (
	"backend/protos/gen/go/auth"
)

type AuthRequest struct {
	ClientID string
	AppID    int
	Login    string
	Password string
}

func AuthRequestFromProto(pb *auth.LoginRequest) (*AuthRequest, error) {
	return &AuthRequest{
		ClientID: pb.ClientId,
		AppID:    int(pb.AppId),
		Login:    pb.Login,
		Password: pb.Password,
	}, nil
}
