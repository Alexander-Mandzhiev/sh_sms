package private_school_models

import "github.com/google/uuid"

func parseStringToUUID(s string) (uuid.UUID, error) {
	if s == "" {
		return uuid.Nil, nil
	}
	id, err := uuid.Parse(s)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}
