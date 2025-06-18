package students_models

import (
	"backend/protos/gen/go/private_school"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Cursor struct {
	ID        uuid.UUID
	CreatedAt time.Time
}

func CursorFromProto(protoCursor *private_school_v1.StudentsCursor) (*Cursor, error) {
	if protoCursor == nil {
		return nil, nil
	}

	id, err := uuid.Parse(protoCursor.Id)
	if err != nil {
		return nil, ErrInvalidCursor
	}

	return &Cursor{
		ID:        id,
		CreatedAt: protoCursor.CreatedAt.AsTime(),
	}, nil
}

func (c *Cursor) ToProto() *private_school_v1.StudentsCursor {
	if c == nil {
		return nil
	}

	return &private_school_v1.StudentsCursor{
		Id:        c.ID.String(),
		CreatedAt: timestamppb.New(c.CreatedAt),
	}
}
