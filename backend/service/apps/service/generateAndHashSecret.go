package service

import (
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) generateAndHashSecret(secretLength int) (string, []byte, error) {
	secret, err := s.generateRandomString(secretLength)
	if err != nil {
		return "", nil, err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
	if err != nil {
		return "", nil, err
	}

	return secret, hashed, nil
}
