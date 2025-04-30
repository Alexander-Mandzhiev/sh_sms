package models

import (
	pb "backend/protos/gen/go/clients/client_types"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type ClientType struct {
	ID          int
	Code        string
	Name        string
	Description string
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func ToProto(ct *ClientType) *pb.ClientType {
	return &pb.ClientType{
		Id:          int32(ct.ID),
		Code:        ct.Code,
		Name:        ct.Name,
		Description: ct.Description,
		IsActive:    ct.IsActive,
		CreatedAt:   timestamppb.New(ct.CreatedAt),
		UpdatedAt:   timestamppb.New(ct.UpdatedAt),
	}
}
