package sso_models

import (
	"backend/protos/gen/go/sso/role_permissions"
	"time"
)

type OperationStatus struct {
	Success   bool      `json:"success"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func OperationStatusFromProto(pb *role_permissions.OperationStatus) *OperationStatus {
	return &OperationStatus{
		Success:   pb.Success,
		Message:   pb.Message,
		Timestamp: pb.Timestamp.AsTime(),
	}
}
