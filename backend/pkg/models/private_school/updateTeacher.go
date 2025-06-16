package private_school_models

import (
	"backend/protos/gen/go/private_school"
	"github.com/google/uuid"
	"strings"
)

type UpdateTeacher struct {
	ID             uuid.UUID `json:"id" db:"id"`
	ClientID       uuid.UUID `json:"client_id" db:"client_id"`
	FullName       *string   `json:"full_name" db:"full_name"`
	Phone          *string   `json:"phone" db:"phone"`
	Email          *string   `json:"email" db:"email"`
	AdditionalInfo *string   `json:"additional_info" db:"additional_info"`
}

func UpdateTeacherFromProto(req *private_school_v1.UpdateTeacherRequest) (*UpdateTeacher, error) {
	teacherID, err := parseStringToUUID(req.GetId())
	if err != nil {
		return nil, ErrInvalidTeacherID
	}

	clientID, err := parseStringToUUID(req.GetClientId())
	if err != nil {
		return nil, ErrInvalidClientID
	}

	update := &UpdateTeacher{
		ID:       teacherID,
		ClientID: clientID,
	}

	if req.FullName != nil {
		val := strings.TrimSpace(*req.FullName)
		update.FullName = &val
	}

	if req.Phone != nil {
		val := cleanPhone(*req.Phone)
		update.Phone = &val
	}

	if req.Email != nil {
		val := strings.ToLower(strings.TrimSpace(*req.Email))
		update.Email = &val
	}

	if req.AdditionalInfo != nil {
		val := strings.TrimSpace(*req.AdditionalInfo)
		update.AdditionalInfo = &val
	}

	if err = update.Validate(); err != nil {
		return nil, err
	}

	return update, nil
}
func (t *UpdateTeacher) ToProto() *private_school_v1.UpdateTeacherRequest {
	return &private_school_v1.UpdateTeacherRequest{
		Id:             t.ID.String(),
		ClientId:       t.ClientID.String(),
		FullName:       t.FullName,
		Phone:          t.Phone,
		Email:          t.Email,
		AdditionalInfo: t.AdditionalInfo,
	}
}

func (t *UpdateTeacher) Validate() error {
	if t.ID == uuid.Nil {
		return ErrInvalidTeacherID
	}
	if t.ClientID == uuid.Nil {
		return ErrInvalidClientID
	}

	if t.FullName != nil {
		if *t.FullName == "" {
			return ErrEmptyFullName
		}
	}
	if t.Phone != nil {
		if *t.Phone != "" && !isValidPhone(*t.Phone) {
			return ErrInvalidPhone
		}
	}

	if t.Email != nil {
		if *t.Email != "" && !isValidEmail(*t.Email) {
			return ErrInvalidEmail
		}
	}

	return nil
}
func (t *UpdateTeacher) Sanitize() {
	if t.FullName != nil {
		val := strings.TrimSpace(*t.FullName)
		t.FullName = &val
	}
	if t.Phone != nil {
		val := cleanPhone(*t.Phone)
		t.Phone = &val
	}
	if t.Email != nil {
		val := strings.ToLower(strings.TrimSpace(*t.Email))
		t.Email = &val
	}
	if t.AdditionalInfo != nil {
		val := strings.TrimSpace(*t.AdditionalInfo)
		t.AdditionalInfo = &val
	}
}
