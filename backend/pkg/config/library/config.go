package config

import (
	"backend/pkg/config/models"
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
	Env        string            `yaml:"env" env:"ENV" env-default:"development"`
	GRPCServer models.GRPCServer `yaml:"grpc_server"`
	DBConfig   DatabaseConfig    `yaml:"database"`
}

type DatabaseConfig struct {
	Postgres models.PostgresConfig `yaml:"postgresql"`
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

	data, err := yaml.Marshal(logCfg)
	if err != nil {
		log.Printf("Failed to marshal config for logging: %v", err)
		return
	}

	log.Printf("Loaded config:\n%s", string(data))
}

func maskPassword(dsn string) string {
	if dsn == "" {
		return ""
	}
	return regexp.MustCompile(`(password=)[^&]+`).ReplaceAllString(dsn, "${1}***")
}
