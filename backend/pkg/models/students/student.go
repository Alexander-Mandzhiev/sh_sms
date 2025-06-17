package students_models

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/private_school"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Student struct {
	ID             uuid.UUID  `json:"id" db:"id"`
	ClientID       uuid.UUID  `json:"client_id" db:"client_id"`
	FullName       string     `json:"full_name" db:"full_name"`
	ContractNumber string     `json:"contract_number" db:"contract_number"`
	Phone          string     `json:"phone" db:"phone"`
	Email          string     `json:"email" db:"email"`
	AdditionalInfo string     `json:"additional_info" db:"additional_info"`
	DeletedAt      *time.Time `json:"deleted_at" db:"deleted_at"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
}

func (s *Student) IsActive() bool {
	return s.DeletedAt == nil
}

func (s *Student) StudentToProto() *private_school_v1.StudentResponse {
	return &private_school_v1.StudentResponse{
		Id:             s.ID.String(),
		ClientId:       s.ClientID.String(),
		FullName:       s.FullName,
		ContractNumber: s.ContractNumber,
		Phone:          s.Phone,
		Email:          s.Email,
		AdditionalInfo: s.AdditionalInfo,
		CreatedAt:      timestamppb.New(s.CreatedAt),
		UpdatedAt:      timestamppb.New(s.UpdatedAt),
		DeletedAt:      utils.TimeToTimestamp(s.DeletedAt),
	}
}

func StudentFromProto(resp *private_school_v1.StudentResponse) (*Student, error) {
	clientID, err := utils.ValidateStringAndReturnUUID(resp.GetClientId())
	if err != nil {
		return nil, ErrInvalidClientID
	}

	studentID, err := utils.ValidateStringAndReturnUUID(resp.GetId())
	if err != nil {
		return nil, ErrInvalidStudentID
	}

	if resp.FullName == "" {
		return nil, ErrEmptyFullName
	}
	if resp.ContractNumber == "" {
		return nil, ErrEmptyContractNumber
	}
	if resp.Phone == "" {
		return nil, ErrEmptyPhone
	}
	if resp.Email == "" {
		return nil, ErrEmptyEmail
	}

	if len(resp.FullName) > MaxFullName {
		return nil, ErrFullNameTooLong
	}
	if len(resp.ContractNumber) > MaxContractNumber {
		return nil, ErrContractNumberTooLong
	}
	if len(resp.Email) > MaxEmail {
		return nil, ErrEmailTooLong
	}

	if !utils.IsValidPhone(resp.Phone) {
		return nil, ErrInvalidPhone
	}
	if !utils.IsValidEmail(resp.Email) {
		return nil, ErrInvalidEmail
	}

	return &Student{
		ID:             studentID,
		ClientID:       clientID,
		FullName:       resp.FullName,
		ContractNumber: resp.ContractNumber,
		Phone:          utils.FormatPhone(resp.Phone),
		Email:          resp.Email,
		AdditionalInfo: resp.AdditionalInfo,
		CreatedAt:      resp.CreatedAt.AsTime(),
		UpdatedAt:      resp.UpdatedAt.AsTime(),
		DeletedAt:      utils.TimestampToTime(resp.DeletedAt),
	}, nil
}
