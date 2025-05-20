package models

import "time"

type GRPCServer struct {
	Address     string        `yaml:"address" env:"GRPC_ADDRESS" env-default:"0.0.0.0"`
	Port        int           `yaml:"port" env:"GRPC_PORT" env-default:"6511"`
	Timeout     time.Duration `yaml:"timeout" env:"GRPC_TIMEOUT" env-default:"10s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"GRPC_IDLE_TIMEOUT" env-default:"60s"`
}
