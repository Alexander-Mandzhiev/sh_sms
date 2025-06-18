package groups_models

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/private_school"
	"strings"

	"github.com/google/uuid"
)

type ListGroupsRequest struct {
	ClientID   uuid.UUID
	PageSize   int32
	Cursor     *int64
	NameFilter *string
}

func ListGroupsParamsFromProto(req *private_school_v1.ListGroupsRequest) (*ListGroupsRequest, error) {
	clientID, err := utils.ValidateStringAndReturnUUID(req.GetClientId())
	if err != nil {
		return nil, ErrInvalidClientID
	}

	pageSize := req.GetPageSize()
	if pageSize < 1 || pageSize > MaxListCount {
		return nil, ErrInvalidPageSize
	}

	params := &ListGroupsRequest{
		ClientID: clientID,
		PageSize: pageSize,
	}

	if req.Cursor != nil {
		cursorVal := req.GetCursor()
		if cursorVal > 0 {
			params.Cursor = &cursorVal
		}
	}

	if req.NameFilter != nil {
		nameFilter := strings.TrimSpace(req.GetNameFilter())
		if nameFilter != "" {
			if len(nameFilter) > MaxFilterValueLength {
				return nil, ErrFilterValueTooLong
			}
			params.NameFilter = &nameFilter
		}
	}

	return params, nil
}

func (lp *ListGroupsRequest) ToProto() *private_school_v1.ListGroupsRequest {
	protoReq := &private_school_v1.ListGroupsRequest{
		ClientId:   lp.ClientID.String(),
		PageSize:   lp.PageSize,
		NameFilter: lp.NameFilter,
	}

	if lp.Cursor != nil {
		cursorVal := *lp.Cursor
		protoReq.Cursor = &cursorVal
	}

	return protoReq
}
