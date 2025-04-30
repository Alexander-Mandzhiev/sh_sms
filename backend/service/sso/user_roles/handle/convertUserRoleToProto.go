package handle

import (
	"backend/protos/gen/go/sso/users_roles"
	"backend/service/sso/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUserRoleToProto(role *models.UserRole) *user_roles.UserRole {
	var expiresAt *timestamppb.Timestamp
	if role.ExpiresAt != nil {
		expiresAt = timestamppb.New(*role.ExpiresAt)
	}

	return &user_roles.UserRole{
		UserId:     role.UserID.String(),
		RoleId:     role.RoleID.String(),
		ClientId:   role.ClientID.String(),
		AppId:      int32(role.AppID),
		AssignedBy: role.AssignedBy.String(),
		ExpiresAt:  expiresAt,
		AssignedAt: timestamppb.New(role.AssignedAt),
	}
}
