package students_models

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/private_school"
	"fmt"
)

type StudentFilter struct {
	FullName       *string
	ContractNumber *string
	Phone          *string
	Email          *string
}

func StudentFilterFromProto(protoFilter *private_school_v1.Filter) (*StudentFilter, error) {
	if protoFilter == nil {
		return nil, nil
	}

	filter := &StudentFilter{
		FullName:       utils.TrimSpacePointer(protoFilter.FullName),
		ContractNumber: utils.TrimSpacePointer(protoFilter.ContractNumber),
		Phone:          utils.TrimSpacePointer(protoFilter.Phone),
		Email:          utils.TrimSpacePointer(protoFilter.Email),
	}

	if filter.Phone != nil && *filter.Phone != "" {
		formatted, err := utils.FormatPhoneOrFail(*filter.Phone)
		if err != nil {
			return nil, ErrInvalidPhone
		}
		filter.Phone = &formatted
	}

	if filter.Email != nil && *filter.Email != "" {
		normalized := utils.NormalizeEmail(*filter.Email)
		filter.Email = &normalized
	}

	if err := validateFilter(filter); err != nil {
		return nil, err
	}

	return filter, nil
}

func validateFilter(filter *StudentFilter) error {
	if filter == nil {
		return nil
	}

	checkLength := func(value *string, maxLen int, fieldName string) error {
		if value != nil && len(*value) > maxLen {
			return fmt.Errorf("%w: %s exceeds max length %d", ErrFilterValueTooLong, fieldName, maxLen)
		}
		return nil
	}

	if err := checkLength(filter.FullName, MaxFullName, "full_name"); err != nil {
		return err
	}
	if err := checkLength(filter.ContractNumber, MaxContractNumber, "contract_number"); err != nil {
		return err
	}
	if err := checkLength(filter.Phone, MaxPhone, "phone"); err != nil {
		return err
	}
	if err := checkLength(filter.Email, MaxEmail, "email"); err != nil {
		return err
	}

	return nil
}

func (sf *StudentFilter) ToProto() *private_school_v1.Filter {
	if sf == nil {
		return nil
	}
	return &private_school_v1.Filter{
		FullName:       sf.FullName,
		ContractNumber: sf.ContractNumber,
		Phone:          sf.Phone,
		Email:          sf.Email,
	}
}
