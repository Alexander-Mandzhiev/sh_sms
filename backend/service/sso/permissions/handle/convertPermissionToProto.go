package handle

import (
	"backend/protos/gen/go/sso/permissions"
	"backend/service/sso/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertPermissionToProto(perm *models.Permission) *permissions.Permission {
	protoPerm := &permissions.Permission{
		Id:          perm.ID.String(),
		Code:        perm.Code,
		Description: perm.Description,
		Category:    perm.Category,
		AppId:       int32(perm.AppID),
		IsActive:    perm.IsActive,
		CreatedAt:   timestamppb.New(perm.CreatedAt),
		UpdatedAt:   timestamppb.New(perm.UpdatedAt),
	}

	if perm.DeletedAt != nil {
		protoPerm.DeletedAt = timestamppb.New(*perm.DeletedAt)
	}

	return protoPerm
}

func convertPermissionsToProto(list []models.Permission) []*permissions.Permission {
	result := make([]*permissions.Permission, len(list))
	for i, p := range list {
		result[i] = convertPermissionToProto(&p)
	}
	return result
}
