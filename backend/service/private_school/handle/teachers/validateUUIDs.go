package teachers_handle

import (
	"errors"
	"github.com/google/uuid"
)

func validateUUIDs(idStr, clientIDStr string) (id, clientID uuid.UUID, err error) {
	if idStr == "" || clientIDStr == "" {
		return uuid.Nil, uuid.Nil, errors.New("teacher_id and client_id are required")
	}

	id, err = uuid.Parse(idStr)
	if err != nil {
		return uuid.Nil, uuid.Nil, errors.New("invalid teacher_id format")
	}

	clientID, err = uuid.Parse(clientIDStr)
	if err != nil {
		return uuid.Nil, uuid.Nil, errors.New("invalid client_id format")
	}

	return id, clientID, nil
}
