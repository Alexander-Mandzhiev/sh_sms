package students_models

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/private_school"
	"github.com/google/uuid"
)

type ListStudentsRequest struct {
	ClientID uuid.UUID
	Count    int32
	Cursor   string
	Filter   *StudentFilter
}

func ListStudentsParamsFromProto(req *private_school_v1.ListStudentsRequest) (*ListStudentsRequest, error) {
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
		cursor := *req.Cursor
		if cursor != "" {
			if _, err := uuid.Parse(cursor); err != nil {
				return nil, ErrInvalidCursor
			}
		}
		params.Cursor = cursor
	}

	if req.Filter != nil {
		filter, err := StudentFilterFromProto(req.GetFilter())
		if err != nil {
			return nil, err
		}
		params.Filter = filter
	}

	return params, nil
}
