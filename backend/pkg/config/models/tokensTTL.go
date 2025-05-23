package models

import "time"

type TokensTTL struct {
	AccessTokenDuration  time.Duration `yaml:"access_token_duration" env:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `yaml:"refresh_token_duration" env:"REFRESH_TOKEN_DURATION"`
}
