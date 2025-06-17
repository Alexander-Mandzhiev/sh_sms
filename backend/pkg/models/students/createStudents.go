package students_models

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/private_school"
	"github.com/google/uuid"
	"strings"
)

type CreateStudent struct {
	ClientID       uuid.UUID `json:"client_id" db:"client_id"`
	FullName       string    `json:"full_name" db:"full_name"`
	ContractNumber string    `json:"contract_number" db:"contract_number"`
	Phone          string    `json:"phone" db:"phone"`
	Email          string    `json:"email" db:"email"`
	AdditionalInfo string    `json:"additional_info" db:"additional_info"`
}

func CreateStudentFromProto(req *private_school_v1.CreateStudentRequest) (*CreateStudent, error) {
	clientID, err := utils.ValidateStringAndReturnUUID(req.GetClientId())
	if err != nil {
		return nil, ErrInvalidClientID
	}

	phone := utils.FormatPhone(req.GetPhone())
	email := utils.NormalizeEmail(req.GetEmail())

	student := &CreateStudent{
		ClientID:       clientID,
		FullName:       strings.TrimSpace(req.GetFullName()),
		ContractNumber: strings.TrimSpace(req.GetContractNumber()),
		Phone:          phone,
		Email:          email,
		AdditionalInfo: req.GetAdditionalInfo(),
	}

	if err = student.Validate(); err != nil {
		return nil, err
	}

	return student, nil
}

func (cs *CreateStudent) ToProto() *private_school_v1.CreateStudentRequest {
	return &private_school_v1.CreateStudentRequest{
		ClientId:       cs.ClientID.String(),
		FullName:       cs.FullName,
		ContractNumber: cs.ContractNumber,
		Phone:          cs.Phone,
		Email:          cs.Email,
		AdditionalInfo: cs.AdditionalInfo,
	}
}

func (cs *CreateStudent) Validate() error {
	if cs.FullName == "" {
		return ErrEmptyFullName
	}
	if cs.ContractNumber == "" {
		return ErrEmptyContractNumber
	}
	if cs.Email == "" {
		return ErrEmptyEmail
	}

	if len(cs.FullName) > MaxFullName {
		return ErrFullNameTooLong
	}
	if len(cs.ContractNumber) > MaxContractNumber {
		return ErrContractNumberTooLong
	}
	if len(cs.Email) > MaxEmail {
		return ErrEmailTooLong
	}

	if !utils.IsValidUUID(cs.ClientID) {
		return ErrInvalidClientID
	}
	if !utils.IsValidPhone(cs.Phone) {
		return ErrInvalidPhone
	}
	if !utils.IsValidEmail(cs.Email) {
		return ErrInvalidEmail
	}

	if len(cs.Phone) < 8 {
		return ErrPhoneTooShort
	}

	return nil
}
