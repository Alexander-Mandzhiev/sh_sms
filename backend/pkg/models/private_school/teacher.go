package private_school_models

import (
	"backend/protos/gen/go/private_school"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Teacher struct {
	ID             string     `db:"id"`
	FullName       string     `db:"full_name"`
	Phone          string     `db:"phone"`
	Email          *string    `db:"email"`
	AdditionalInfo *string    `db:"additional_info"`
	DeletedAt      *time.Time `db:"deleted_at"`
	CreatedAt      time.Time  `db:"created_at"`
	UpdatedAt      time.Time  `db:"updated_at"`
}

func (t *Teacher) IsActive() bool {
	return t.DeletedAt == nil
}

func TeacherToProto(teacher *Teacher) *private_school_v1.TeacherResponse {
	response := &private_school_v1.TeacherResponse{
		Id:        teacher.ID,
		FullName:  teacher.FullName,
		Phone:     teacher.Phone,
		CreatedAt: timestamppb.New(teacher.CreatedAt),
		UpdatedAt: timestamppb.New(teacher.UpdatedAt),
	}

	if teacher.Email != nil {
		response.Email = *teacher.Email
	}

	if teacher.AdditionalInfo != nil {
		response.AdditionalInfo = *teacher.AdditionalInfo
	}

	if teacher.DeletedAt != nil {
		response.DeletedAt = timestamppb.New(*teacher.DeletedAt)
	}

	return response
}

func TeacherResponseToTeacher(resp *private_school_v1.TeacherResponse) *Teacher {
	teacher := &Teacher{
		ID:        resp.Id,
		FullName:  resp.FullName,
		Phone:     resp.Phone,
		CreatedAt: resp.CreatedAt.AsTime(),
		UpdatedAt: resp.UpdatedAt.AsTime(),
	}

	if resp.Email != "" {
		teacher.Email = &resp.Email
	}

	if resp.AdditionalInfo != "" {
		teacher.AdditionalInfo = &resp.AdditionalInfo
	}

	if resp.DeletedAt != nil {
		deletedAt := resp.DeletedAt.AsTime()
		teacher.DeletedAt = &deletedAt
	}

	return teacher
}
