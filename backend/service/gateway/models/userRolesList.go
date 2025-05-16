package models

import "backend/protos/gen/go/sso/users_roles"

type UserRolesList struct {
	Assignments []UserRole `json:"assignments"`
	TotalCount  int32      `json:"total_count"`
	CurrentPage int32      `json:"current_page"`
	AppID       int32      `json:"app_id"`
}

func UserRolesListFromProto(pb *user_roles.UserRolesResponse) *UserRolesList {
	list := &UserRolesList{
		TotalCount:  pb.TotalCount,
		CurrentPage: pb.CurrentPage,
		AppID:       pb.AppId,
	}

	for _, role := range pb.Assignments {
		list.Assignments = append(list.Assignments, *UserRoleFromProto(role))
	}
	return list
}
