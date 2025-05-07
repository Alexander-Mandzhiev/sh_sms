package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Env        string         `yaml:"env" env:"ENV" env-default:"development"`
	HTTPServer HTTPServer     `yaml:"http_server"`
	Services   ServicesConfig `yaml:"services"`
	Frontend   Frontend       `yaml:"frontend"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env:"ADDRESS" env-default:"0.0.0.0"`       // Адрес сервера
	Port        int           `yaml:"port" env:"PORT" env-default:"6000"`                // Порт сервера
	Timeout     time.Duration `yaml:"timeout" env:"TIMEOUT" env-default:"5s"`            // Таймаут запроса
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"IDLE_TIMEOUT" env-default:"60s"` // Idle таймаут
}

type ServicesConfig struct {
}

type Frontend struct {
	Addr string `yaml:"addr" env:"FRONTEND_ADDR" env-default:"0.0.0.0:6900"`
}

func MustLoad() *Config {
	configPath := fetchConfigFlag()
	if configPath == "" {
		return loadingDataInEnv()
	}
	return MustLoadByPath(configPath)
}

func MustLoadByPath(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}
	return &cfg
}

func fetchConfigFlag() string {
	var res string
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}

func loadingDataInEnv() *Config {
	loadEnv()

	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil || port <= 0 {
		log.Printf("Warning: Invalid PORT value in environment variables, using default value %d.", 6000)
		port = 6000
	}

	return &Config{
		HTTPServer: HTTPServer{
			Address:     os.Getenv("ADDRESS"),
			Port:        port,
			Timeout:     parseDuration(os.Getenv("TIMEOUT"), 5*time.Second),
			IdleTimeout: parseDuration(os.Getenv("IDLE_TIMEOUT"), 60*time.Second),
		},
		Services: ServicesConfig{},
		Frontend: Frontend{
			Addr: os.Getenv("FRONTEND_ADDR"),
		},
	}
}

func loadEnv() {
	if err := godotenv.Load(".gateway.env"); err != nil {
		log.Println("Warning: .gateway.env file not found, using default values.")
	}
}

func parseDuration(value string, defaultValue time.Duration) time.Duration {
	duration, err := time.ParseDuration(value)
	if err != nil || duration <= 0 {
		log.Printf("Warning: Invalid TIMEOUT or IDLE_TIMEOUT value in environment variables, using default value %v.", defaultValue)
		return defaultValue
	}
	return duration
}
