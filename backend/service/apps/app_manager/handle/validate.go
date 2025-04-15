package handle

import (
	pb "backend/protos/gen/go/apps/app_manager"
	"fmt"
)

func validateName(name string, maxLength int) error {
	if name == "" {
		return ErrEmptyName
	}
	if len(name) > maxLength {
		return fmt.Errorf("%w: max length %d", ErrInvalidName, maxLength)
	}
	return nil
}

func validateCode(code string, maxLength int) error {
	if code == "" {
		return ErrEmptyCode
	}
	if len(code) > 50 {
		return fmt.Errorf("%w: max length %d", ErrInvalidCode, maxLength)
	}
	return nil
}

func validateUpdateRequest(req *pb.UpdateRequest) error {
	if err := validateID(req.GetId()); err != nil {
		return err
	}

	if req.Code == nil && req.Name == nil && req.Description == nil && req.IsActive == nil {
		return ErrNoUpdateFields
	}

	if req.Name != nil {
		if err := validateName(req.GetName(), 250); err != nil {
			return err
		}
	}

	if req.Code != nil {
		if err := validateCode(req.GetCode(), 50); err != nil {
			return err
		}
	}

	return nil
}

func validateNoConflict(req *pb.AppIdentifier) error {
	hasBoth := req.GetId() != 0 && req.GetCode() != ""
	if hasBoth {
		return ErrConflictParams
	}
	return nil
}

func validateAtLeastOne(req *pb.AppIdentifier) error {
	hasAny := req.GetId() != 0 || req.GetCode() != ""
	if !hasAny {
		return ErrIdentifierRequired
	}
	return nil
}

func validateID(id int32) error {
	if id <= 0 {
		return fmt.Errorf("%w: %d", ErrInvalidID, id)
	}
	return nil
}
