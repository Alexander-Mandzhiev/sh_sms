package teachers_models

import (
	"backend/protos/gen/go/private_school"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type ListTeachersFilter struct {
	Cursor         *Cursor
	Limit          int32
	ClientID       uuid.UUID
	IncludeDeleted bool
}

func ListTeachersFilterFromProto(req *private_school_v1.ListTeachersRequest) (*ListTeachersFilter, error) {
	if req.GetClientId() == "" {
		return nil, errors.New("client_id is required")
	}

	clientID, err := uuid.Parse(req.GetClientId())
	if err != nil {
		return nil, fmt.Errorf("invalid client_id format: %w", err)
	}

	limit := req.GetLimit()
	if limit <= 0 || limit > 1000 {
		limit = 50
	}

	return &ListTeachersFilter{
		Cursor:         CursorFromProto(req.Cursor),
		Limit:          limit,
		ClientID:       clientID,
		IncludeDeleted: req.IncludeDeleted,
	}, nil
}

func ListTeachersFilterToProto(filter *ListTeachersFilter) *private_school_v1.ListTeachersRequest {
	return &private_school_v1.ListTeachersRequest{
		ClientId:       filter.ClientID.String(),
		Cursor:         CursorToProto(filter.Cursor),
		Limit:          filter.Limit,
		IncludeDeleted: filter.IncludeDeleted,
	}
}

func (f *ListTeachersFilter) Validate() error {
	switch {
	case f.ClientID == uuid.Nil:
		return ErrInvalidClientID
	case f.Limit <= 0 || f.Limit > 1000:
		return ErrInvalidLimitRange
	}
	return nil
}
