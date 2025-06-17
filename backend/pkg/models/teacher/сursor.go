package teachers_models

import (
	"backend/protos/gen/go/private_school"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Cursor struct {
	LastID    string
	CreatedAt time.Time
}

func CursorFromProto(pb *private_school_v1.Cursor) *Cursor {
	if pb == nil {
		return nil
	}
	return &Cursor{
		LastID:    pb.LastId,
		CreatedAt: pb.CreatedAt.AsTime(),
	}
}

func CursorToProto(cursor *Cursor) *private_school_v1.Cursor {
	if cursor == nil {
		return nil
	}
	return &private_school_v1.Cursor{
		LastId:    cursor.LastID,
		CreatedAt: timestamppb.New(cursor.CreatedAt),
	}
}

func EncodeCursorForURL(cursor *Cursor) (string, error) {
	if cursor == nil {
		return "", nil
	}

	type cursorData struct {
		LastID    string `json:"last_id"`
		CreatedAt string `json:"created_at"`
	}

	data := cursorData{
		LastID:    cursor.LastID,
		CreatedAt: cursor.CreatedAt.Format(time.RFC3339Nano),
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal cursor: %w", err)
	}

	return base64.URLEncoding.EncodeToString(jsonData), nil
}

func DecodeCursorFromURL(cursorStr string) (*Cursor, error) {
	if cursorStr == "" {
		return nil, nil
	}

	jsonData, err := base64.URLEncoding.DecodeString(cursorStr)
	if err != nil {
		return nil, fmt.Errorf("failed to decode cursor: %w", err)
	}

	var data struct {
		LastID    string `json:"last_id"`
		CreatedAt string `json:"created_at"`
	}

	if err = json.Unmarshal(jsonData, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal cursor: %w", err)
	}

	if data.LastID == "" || data.CreatedAt == "" {
		return nil, errors.New("invalid cursor data")
	}

	var createdAt time.Time
	if data.CreatedAt != "" {
		createdAt, err = time.Parse(time.RFC3339Nano, data.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to parse cursor time: %w", err)
		}
	} else {
		createdAt = time.Time{}
	}

	return &Cursor{
		LastID:    data.LastID,
		CreatedAt: createdAt,
	}, nil

	return &Cursor{
		LastID:    data.LastID,
		CreatedAt: createdAt,
	}, nil
}
