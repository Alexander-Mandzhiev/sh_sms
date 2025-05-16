package service

import "context"

func (s *AuthService) Logout(ctx context.Context, accessToken, refreshToken string) error {
	return nil
}
