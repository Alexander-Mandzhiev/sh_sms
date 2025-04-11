package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Env        string         `yaml:"env" env:"ENV" env-default:"development"`
	GRPCServer GRPCServer     `yaml:"grpc_server"`
	DBConfig   DatabaseConfig `yaml:"database"`
}

type GRPCServer struct {
	Address     string        `yaml:"address" env:"GRPC_ADDRESS" env-default:"0.0.0.0"`
	Port        int           `yaml:"port" env:"GRPC_PORT" env-default:"6511"`
	Timeout     time.Duration `yaml:"timeout" env:"GRPC_TIMEOUT" env-default:"10s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"GRPC_IDLE_TIMEOUT" env-default:"60s"`
}

type DatabaseConfig struct {
	Postgres PostgresConfig `yaml:"postgresql"`
}

type PostgresConfig struct {
	ConnectionString   string        `yaml:"connection_string" env:"POSTGRES_CONNECTION_STRING" env-required:"true"`
	MaxOpenConnections int32         `yaml:"max_open_connections" env:"POSTGRES_MAX_OPEN_CONNS" env-default:"20"`
	MaxIdleConnections int32         `yaml:"max_idle_connections" env:"POSTGRES_MAX_IDLE_CONNS" env-default:"10"`
	ConnMaxLifetime    time.Duration `yaml:"conn_max_lifetime" env:"POSTGRES_CONN_MAX_LIFETIME" env-default:"30m"`
	ConnMaxIdleTime    time.Duration `yaml:"conn_max_idle_time" env:"POSTGRES_CONN_MAX_IDLE_TIME" env-default:"5m"`
	HealthCheckPeriod  time.Duration `yaml:"health_check_period" env:"POSTGRES_HEALTH_CHECK_PERIOD" env-default:"1m"`
	ConnectTimeout     time.Duration `yaml:"connect_timeout" env:"POSTGRES_CONNECT_TIMEOUT" env-default:"5s"`
}

func (c *Config) Validate() error {
	validEnvs := map[string]bool{
		"development": true,
		"production":  true,
		"test":        true,
		"staging":     true,
	}

	if !validEnvs[c.Env] {
		return fmt.Errorf("invalid env value: %s, expected one of: development, production, test, staging", c.Env)
	}

	if c.GRPCServer.Port <= 0 || c.GRPCServer.Port > 65535 {
		return fmt.Errorf("invalid port: %d, must be between 1 and 65535", c.GRPCServer.Port)
	}

	if c.DBConfig.Postgres.MaxOpenConnections < c.DBConfig.Postgres.MaxIdleConnections {
		return fmt.Errorf("max_open_connections (%d) must be >= max_idle_connections (%d)",
			c.DBConfig.Postgres.MaxOpenConnections, c.DBConfig.Postgres.MaxIdleConnections)
	}

	return nil
}

func Initialize() *Config {
	cfg := MustLoad()
	if err := cfg.Validate(); err != nil {
		log.Fatalf("Invalid configuration: %v", err)
	}
	logConfig(cfg)
	return cfg
}

func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath != "" {
		return loadFromYAML(configPath)
	}
	return loadFromEnv()
}

func loadFromYAML(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configPath)
	}

	file, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var cfg Config
	if err = yaml.Unmarshal(file, &cfg); err != nil {
		log.Fatalf("Failed to parse YAML config: %v", err)
	}

	return &cfg
}

func loadFromEnv() *Config {
	loadEnvFiles()

	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Failed to read environment variables: %v", err)
	}

	return &cfg
}

func fetchConfigPath() string {
	var path string
	flag.StringVar(&path, "config", "", "path to config file")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}

	return path
}

func loadEnvFiles() {
	envFiles := []string{".env", ".env.local", "config/.env", "config/.env.local"}
	for _, file := range envFiles {
		if err := godotenv.Load(file); err == nil {
			log.Printf("Loaded environment variables from %s", file)
			return
		}
	}
	log.Println("No .env file found, using default values and environment variables")
}

func logConfig(cfg *Config) {
	logCfg := *cfg
	logCfg.DBConfig.Postgres.ConnectionString = maskPassword(logCfg.DBConfig.Postgres.ConnectionString)

	_, err := yaml.Marshal(logCfg)
	if err != nil {
		log.Printf("Failed to marshal config for logging: %v", err)
		return
	}
}

func maskPassword(dsn string) string {
	if dsn == "" {
		return ""
	}
	return regexp.MustCompile(`(password=)[^&]+`).ReplaceAllString(dsn, "${1}***")
}
