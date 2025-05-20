package config

import (
	"backend/pkg/config/models"
	"errors"
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"regexp"
)

type Config struct {
	ServiceName string            `yaml:"service_name" env:"SERVICE_NAME" env-default:"auth-service"`
	Env         string            `yaml:"env" env:"ENV" env-default:"development"`
	GRPCServer  models.GRPCServer `yaml:"grpc_server"`
	DBConfig    DatabaseConfig    `yaml:"database"`
	Services    ServicesConfig    `yaml:"services"`
	JWT         models.JWTConfig  `yaml:"jwt"`
}

type ServicesConfig struct {
	AppsAddr    string `yaml:"apps_addr" env:"APPS_ADDR" env-required:"true"`
	ClientsAddr string `yaml:"clients_addr" env:"CLIENTS_ADDR" env-required:"true"`
	SSOAddr     string `yaml:"sso_addr" env:"SSO_ADDR" env-required:"true"`
}

type DatabaseConfig struct {
	Postgres models.PostgresConfig `yaml:"postgresql"`
}

func (c *Config) Validate() error {
	validEnvs := map[string]bool{"development": true, "production": true, "test": true, "staging": true}
	if !validEnvs[c.Env] {
		return fmt.Errorf("invalid env: %s", c.Env)
	}

	if c.GRPCServer.Port <= 0 || c.GRPCServer.Port > 65535 {
		return fmt.Errorf("invalid port: %d", c.GRPCServer.Port)
	}

	if c.DBConfig.Postgres.MaxOpenConnections < c.DBConfig.Postgres.MaxIdleConnections {
		return fmt.Errorf("invalid connection pool: max_open(%d) < max_idle(%d)",
			c.DBConfig.Postgres.MaxOpenConnections, c.DBConfig.Postgres.MaxIdleConnections)
	}

	if c.Services.AppsAddr == "" {
		return errors.New("apps service address is required")
	}

	if c.JWT.AccessDuration == 0 || c.JWT.RefreshDuration == 0 {
		return errors.New("JWT durations must be set")
	}

	if len(c.JWT.Audiences) == 0 {
		return errors.New("JWT audiences are required")
	}

	return nil
}

func Initialize() *Config {
	cfg := MustLoad()
	if err := cfg.Validate(); err != nil {
		log.Fatalf("Configuration error: %v", err)
	}
	logConfig(cfg)
	return cfg
}

func MustLoad() *Config {
	if path := fetchConfigPath(); path != "" {
		return loadFromYAML(path)
	}
	return loadFromEnv()
}

func loadFromYAML(path string) *Config {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("Config file not found: %s", path)
	}

	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	var cfg Config
	if err = yaml.Unmarshal(file, &cfg); err != nil {
		log.Fatalf("Config parsing failed: %v", err)
	}

	return &cfg
}

func loadFromEnv() *Config {
	loadEnvFiles()

	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Env vars loading failed: %v", err)
	}

	return &cfg
}

func fetchConfigPath() string {
	var path string
	flag.StringVar(&path, "config", "", "Path to config file")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}

	return path
}

func loadEnvFiles() {
	envFiles := []string{".env", ".env.local"}
	for _, file := range envFiles {
		if err := godotenv.Load(file); err == nil {
			log.Printf("Loaded env from %s", file)
			return
		}
	}
	log.Println("No .env files found")
}

func logConfig(cfg *Config) {
	logCfg := *cfg
	logCfg.DBConfig.Postgres.ConnectionString = maskPassword(logCfg.DBConfig.Postgres.ConnectionString)

	data, err := yaml.Marshal(logCfg)
	if err != nil {
		log.Printf("Failed to marshal config: %v", err)
		return
	}

	log.Printf("Active configuration:\n---\n%s\n---", string(data))
}

func maskPassword(dsn string) string {
	if dsn == "" {
		return ""
	}
	return regexp.MustCompile(`(password=)[^&]+`).ReplaceAllString(dsn, "${1}***")
}
