package students_models

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/private_school"
	"github.com/google/uuid"
)

type ListStudentsRequest struct {
	ClientID uuid.UUID
	Count    int32
	Cursor   *Cursor
	Filter   string
}

func ListStudentsRequestFromProto(req *private_school_v1.ListStudentsRequest) (*ListStudentsRequest, error) {
	clientID, err := utils.ValidateStringAndReturnUUID(req.GetClientId())
	if err != nil {
		return nil, ErrInvalidClientID
	}

	params := &ListStudentsRequest{
		ClientID: clientID,
		Count:    DefaultListCount,
	}

	if req.Count != nil {
		count := *req.Count
		if count < 1 || count > MaxListCount {
			return nil, ErrInvalidCountValue
		}
		params.Count = count
	}

	if req.Cursor != nil {
		var cursor *Cursor
		cursor, err = CursorFromProto(req.Cursor)
		if err != nil {
			return nil, err
		}
		params.Cursor = cursor
	}

	if req.Filter != nil {
		if len(params.Filter) > MaxFilterValueLength {
			return nil, ErrFilterValueTooLong
		}
		params.Filter = *req.Filter
	}

	return params, nil
}
