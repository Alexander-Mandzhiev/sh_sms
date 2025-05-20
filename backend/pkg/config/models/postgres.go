package models

import "time"

type PostgresConfig struct {
	ConnectionString   string        `yaml:"connection_string" env:"POSTGRES_CONNECTION_STRING" env-required:"true"`
	MaxOpenConnections int32         `yaml:"max_open_connections" env:"POSTGRES_MAX_OPEN_CONNS" env-default:"20"`
	MaxIdleConnections int32         `yaml:"max_idle_connections" env:"POSTGRES_MAX_IDLE_CONNS" env-default:"10"`
	ConnMaxLifetime    time.Duration `yaml:"conn_max_lifetime" env:"POSTGRES_CONN_MAX_LIFETIME" env-default:"30m"`
	ConnMaxIdleTime    time.Duration `yaml:"conn_max_idle_time" env:"POSTGRES_CONN_MAX_IDLE_TIME" env-default:"5m"`
	HealthCheckPeriod  time.Duration `yaml:"health_check_period" env:"POSTGRES_HEALTH_CHECK_PERIOD" env-default:"1m"`
	ConnectTimeout     time.Duration `yaml:"connect_timeout" env:"POSTGRES_CONNECT_TIMEOUT" env-default:"5s"`
}
