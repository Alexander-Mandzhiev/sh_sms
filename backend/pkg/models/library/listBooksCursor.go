package library_models

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"time"
)

type ListBooksCursor struct {
	LastID    int64     `json:"last_id"`
	CreatedAt time.Time `json:"created_at"`
}

func NewCursor(lastID int64, createdAt time.Time) *ListBooksCursor {
	return &ListBooksCursor{
		LastID:    lastID,
		CreatedAt: createdAt,
	}
}

func (c *ListBooksCursor) Encode() (string, error) {
	jsonData, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(jsonData), nil
}

func decodeCursor(token string) (*ListBooksCursor, error) {
	if token == "" {
		return nil, errors.New("empty token")
	}

	jsonData, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}

	var cursor ListBooksCursor
	if err = json.Unmarshal(jsonData, &cursor); err != nil {
		return nil, err
	}

	return &cursor, nil
}
