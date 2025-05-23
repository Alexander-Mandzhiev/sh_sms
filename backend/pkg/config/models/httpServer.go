package models

import "time"

type HTTPServer struct {
	Address     string        `yaml:"address" env:"ADDRESS" env-default:"0.0.0.0"`
	Port        int           `yaml:"port" env:"PORT" env-default:"6000"`
	Timeout     time.Duration `yaml:"timeout" env:"TIMEOUT" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"IDLE_TIMEOUT" env-default:"60s"`
}
