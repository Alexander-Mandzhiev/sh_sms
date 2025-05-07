package models

import (
	pb "backend/protos/gen/go/clients/clients"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Client struct {
	ID          string
	Name        string
	Description string
	TypeID      int
	Website     string
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (c *Client) ToProto() *pb.Client {
	return &pb.Client{
		Id:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		TypeId:      int32(c.TypeID),
		Website:     c.Website,
		IsActive:    c.IsActive,
		CreatedAt:   timestamppb.New(c.CreatedAt),
		UpdatedAt:   timestamppb.New(c.UpdatedAt),
	}
}
