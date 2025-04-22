package handle

import (
	"backend/protos/gen/go/sso/roles"
	"backend/service/sso/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertRoleToProto(role *models.Role) *roles.Role {
	protoRole := &roles.Role{
		Id:            role.ID.String(),
		ClientId:      role.ClientID.String(),
		Name:          role.Name,
		Description:   role.Description,
		Level:         int32(role.Level),
		IsCustom:      role.IsCustom,
		IsActive:      role.DeletedAt == nil,
		CreatedAt:     timestamppb.New(role.CreatedAt),
		DeletedAt:     nil,
		PermissionIds: extractPermissionIDs(role.Permissions),
	}

	if role.ParentRoleID != nil {
		parentID := role.ParentRoleID.String()
		protoRole.ParentRoleId = &parentID
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

func extractPermissionIDs(perms []models.Permission) []string {
	ids := make([]string, 0, len(perms))
	for _, p := range perms {
		ids = append(ids, p.ID.String())
	}
	return ids
}
