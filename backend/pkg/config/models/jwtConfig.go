package models

import "time"

type JWTConfig struct {
	AccessDuration  time.Duration `yaml:"access_duration" env:"JWT_ACCESS_DURATION" env-default:"15m"`
	RefreshDuration time.Duration `yaml:"refresh_duration" env:"JWT_REFRESH_DURATION" env-default:"168h"`
	Audiences       []string      `yaml:"audiences" env:"JWT_AUDIENCES" env-separator:"," env-default:"web-app,mobile-api"`
}
