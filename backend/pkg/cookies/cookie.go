package cookies

import (
	"net/http"
	"time"
)

type Config struct {
	Secure   bool          // Использовать Secure (HTTPS only)
	SameSite http.SameSite // Политика SameSite
	Domain   string        // Домен для cookie
	Path     string        // Путь для cookie
	HTTPOnly bool          // Запретить доступ через JavaScript
}

var DefaultConfig = Config{
	Secure:   true,
	SameSite: http.SameSiteLaxMode,
	Path:     "/",
	HTTPOnly: true,
}

func SetRefreshCookie(w http.ResponseWriter, token string, cfg Config, ttl time.Duration) {
	expires := time.Now().Add(ttl)

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    token,
		Path:     cfg.Path,
		Domain:   cfg.Domain,
		Secure:   cfg.Secure,
		HttpOnly: cfg.HTTPOnly,
		SameSite: cfg.SameSite,
		Expires:  expires,
		MaxAge:   int(ttl.Seconds()),
	})
}

// RemoveRefreshCookie удаляет refresh token cookie
func RemoveRefreshCookie(w http.ResponseWriter, cfg Config) {
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     cfg.Path,
		Domain:   cfg.Domain,
		Secure:   cfg.Secure,
		HttpOnly: cfg.HTTPOnly,
		SameSite: cfg.SameSite,
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
	})
}
