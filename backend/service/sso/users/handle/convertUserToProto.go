package handle

import (
	"backend/protos/gen/go/sso/users"
	"backend/service/sso/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUserToProto(u models.User) *users.User {
	var deletedAt *timestamppb.Timestamp
	if u.DeletedAt != nil {
		deletedAt = timestamppb.New(*u.DeletedAt)
	}

	return &users.User{
		Id:        u.ID.String(),
		ClientId:  u.ClientID.String(),
		Email:     u.Email,
		FullName:  u.FullName,
		Phone:     u.Phone,
		IsActive:  u.IsActive,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
		DeletedAt: deletedAt,
	}
}
