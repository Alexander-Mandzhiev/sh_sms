package teachers_models

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/private_school"
	"github.com/google/uuid"
	"strings"
)

type CreateTeacher struct {
	ID             uuid.UUID `json:"id" db:"id"`
	ClientID       uuid.UUID `json:"client_id" db:"client_id"`
	FullName       string    `json:"full_name" db:"full_name"`
	Phone          string    `json:"phone" db:"phone"`
	Email          *string   `json:"email" db:"email"`
	AdditionalInfo *string   `json:"additional_info" db:"additional_info"`
}

func CreateTeacherFromProto(req *private_school_v1.CreateTeacherRequest) (*CreateTeacher, error) {
	teacherID, err := parseStringToUUID(req.GetId())
	if err != nil {
		return nil, ErrInvalidTeacherID
	}

	clientID, err := parseStringToUUID(req.GetClientId())
	if err != nil {
		return nil, ErrInvalidClientID
	}

	phone := utils.CleanPhone(req.GetPhone())

	teacher := &CreateTeacher{
		ID:             teacherID,
		ClientID:       clientID,
		FullName:       strings.TrimSpace(req.GetFullName()),
		Phone:          phone,
		Email:          req.Email,
		AdditionalInfo: req.AdditionalInfo,
	}

	if err = teacher.Validate(); err != nil {
		return nil, err
	}

	return teacher, nil
}

func (t *CreateTeacher) ToProto() *private_school_v1.CreateTeacherRequest {
	return &private_school_v1.CreateTeacherRequest{
		Id:             t.ID.String(),
		ClientId:       t.ClientID.String(),
		FullName:       t.FullName,
		Phone:          t.Phone,
		Email:          t.Email,
		AdditionalInfo: t.AdditionalInfo,
	}
}

func (t *CreateTeacher) Validate() error {
	switch {
	case t.FullName == "":
		return ErrEmptyFullName
	case t.ID == uuid.Nil:
		return ErrInvalidTeacherID
	case t.ClientID == uuid.Nil:
		return ErrInvalidClientID
	case t.Phone != "" && !utils.IsValidPhone(t.Phone):
		return ErrInvalidPhone
	case t.Email != nil && *t.Email != "" && !utils.IsValidEmail(*t.Email):
		return ErrInvalidEmail
	}
	return nil
}
