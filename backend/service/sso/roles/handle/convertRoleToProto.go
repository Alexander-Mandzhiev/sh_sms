package handle

import (
	"backend/protos/gen/go/sso/roles"
	"backend/service/sso/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertRoleToProto(role *models.Role) *roles.Role {
	protoRole := &roles.Role{
		Id:          role.ID.String(),
		ClientId:    role.ClientID.String(),
		AppId:       int32(role.AppID),
		Name:        role.Name,
		Description: &role.Description,
		Level:       int32(role.Level),
		IsCustom:    role.IsCustom,
		IsActive:    role.IsActive,
		CreatedAt:   timestamppb.New(role.CreatedAt),
		UpdatedAt:   timestamppb.New(role.UpdatedAt),
	}

	if role.CreatedBy != nil {
		createdBy := role.CreatedBy.String()
		protoRole.CreatedBy = &createdBy
	}

	if role.DeletedAt != nil {
		protoRole.DeletedAt = timestamppb.New(*role.DeletedAt)
	}

	return protoRole
}
