package models

type FileStorageConfig struct {
	BaseDir       string `yaml:"base_dir" env:"FILE_STORAGE_BASE_DIR" env-default:"./uploads"`
	MaxFileSizeMB int    `yaml:"max_file_size_mb" env:"FILE_STORAGE_MAX_SIZE_MB" env-default:"50"`
	BaseURL       string `yaml:"base_url" env:"FILE_STORAGE_BASE_URL" env-default:"http://localhost:8080/files"`
}
