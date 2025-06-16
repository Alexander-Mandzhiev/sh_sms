package private_school_models

import (
	"backend/protos/gen/go/private_school"
	"github.com/google/uuid"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Teacher struct {
	ID             uuid.UUID  `json:"id" db:"id"`
	ClientID       uuid.UUID  `json:"client_id" db:"client_id"`
	FullName       string     `json:"full_name" db:"full_name"`
	Phone          string     `json:"phone" db:"phone"`
	Email          *string    `json:"email" db:"email"`
	AdditionalInfo *string    `json:"additional_info" db:"additional_info"`
	DeletedAt      *time.Time `json:"deleted_at" db:"deleted_at"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
}

func (t *Teacher) IsActive() bool {
	return t.DeletedAt == nil
}

func (t *Teacher) TeacherToProto() *private_school_v1.TeacherResponse {
	response := &private_school_v1.TeacherResponse{
		Id:        t.ID.String(),
		ClientId:  t.ClientID.String(),
		FullName:  t.FullName,
		Phone:     t.Phone,
		CreatedAt: timestamppb.New(t.CreatedAt),
		UpdatedAt: timestamppb.New(t.UpdatedAt),
	}

	if t.Email != nil {
		response.Email = *t.Email
	}

	if t.AdditionalInfo != nil {
		response.AdditionalInfo = *t.AdditionalInfo
	}

	if t.DeletedAt != nil {
		response.DeletedAt = timestamppb.New(*t.DeletedAt)
	}

	return response
}

func TeacherFromProto(resp *private_school_v1.TeacherResponse) (*Teacher, error) {
	clientID, err := parseStringToUUID(resp.GetClientId())
	if err != nil {
		return nil, ErrInvalidClientID
	}
	teacherId, err := parseStringToUUID(resp.GetId())
	if err != nil {
		return nil, ErrInvalidTeacherID
	}

	teacher := &Teacher{
		ID:        teacherId,
		ClientID:  clientID,
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

	return teacher, nil
}
