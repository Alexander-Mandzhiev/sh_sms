package students_models

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/private_school"
	"github.com/google/uuid"
)

type UpdateStudent struct {
	ID             uuid.UUID
	ClientID       uuid.UUID
	FullName       *string
	ContractNumber *string
	Phone          *string
	Email          *string
	AdditionalInfo *string
}

func UpdateStudentFromProto(req *private_school_v1.UpdateStudentRequest) (*UpdateStudent, error) {
	studentID, err := utils.ValidateStringAndReturnUUID(req.GetId())
	if err != nil {
		return nil, ErrInvalidStudentID
	}

	clientID, err := utils.ValidateStringAndReturnUUID(req.GetClientId())
	if err != nil {
		return nil, ErrInvalidClientID
	}

	update := &UpdateStudent{
		ID:             studentID,
		ClientID:       clientID,
		FullName:       nil,
		ContractNumber: nil,
		Phone:          nil,
		Email:          nil,
		AdditionalInfo: nil,
	}

	if req.FullName != nil {
		val := utils.TrimSpacePointer(req.FullName)
		update.FullName = val
	}
	if req.ContractNumber != nil {
		val := utils.TrimSpacePointer(req.ContractNumber)
		update.ContractNumber = val
	}
	if req.Phone != nil {
		phone := utils.TrimSpacePointer(req.Phone)
		val := utils.FormatPhone(*phone)
		update.Phone = &val
	}
	if req.Email != nil {
		email := utils.TrimSpacePointer(req.Email)
		val := utils.NormalizeEmail(*email)
		update.Email = &val
	}
	if req.AdditionalInfo != nil {
		update.AdditionalInfo = req.AdditionalInfo
	}

	if err = update.Validate(); err != nil {
		return nil, err
	}

	return update, nil
}

func (us *UpdateStudent) ToProto() *private_school_v1.UpdateStudentRequest {
	return &private_school_v1.UpdateStudentRequest{
		Id:             us.ID.String(),
		ClientId:       us.ClientID.String(),
		FullName:       us.FullName,
		ContractNumber: us.ContractNumber,
		Phone:          us.Phone,
		Email:          us.Email,
		AdditionalInfo: us.AdditionalInfo,
	}
}

func (us *UpdateStudent) Validate() error {
	if !utils.IsValidUUID(us.ID) {
		return ErrInvalidStudentID
	}
	if !utils.IsValidUUID(us.ClientID) {
		return ErrInvalidClientID
	}

	if us.FullName != nil {
		if *us.FullName == "" {
			return ErrEmptyFullName
		}
		if len(*us.FullName) > MaxFullName {
			return ErrFullNameTooLong
		}
	}

	if us.ContractNumber != nil {
		if *us.ContractNumber == "" {
			return ErrEmptyContractNumber
		}
		if len(*us.ContractNumber) > MaxContractNumber {
			return ErrContractNumberTooLong
		}
	}

	if us.Phone != nil {
		if *us.Phone == "" {
			return ErrEmptyPhone
		}
		if !utils.IsValidPhone(*us.Phone) {
			return ErrInvalidPhone
		}
		if len(*us.Phone) < MinPhoneLength {
			return ErrPhoneTooShort
		}
		if len(*us.Phone) > MaxPhone {
			return ErrPhoneTooLong
		}
	}

	if us.Email != nil {
		if *us.Email == "" {
			return ErrEmptyEmail
		}
		if len(*us.Email) > MaxEmail {
			return ErrEmailTooLong
		}
		if !utils.IsValidEmail(*us.Email) {
			return ErrInvalidEmail
		}
	}

	return nil
}
