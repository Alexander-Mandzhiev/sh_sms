package library_models

type UploadedFile struct {
	FilePath string // Относительный путь
	Size     int64
	Checksum string // SHA256 хеш
}
