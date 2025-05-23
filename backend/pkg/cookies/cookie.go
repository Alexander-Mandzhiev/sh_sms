package cookies

import (
	"net/http"
	"time"
)

type Config struct {
	Secure   bool
	SameSite http.SameSite
	Domain   string
	Path     string
	HTTPOnly bool
}

var DefaultConfig = Config{
	Secure:   false,
	SameSite: http.SameSiteNoneMode,
	Domain:   "localhost",
	Path:     "/",
	HTTPOnly: true,
}

func SetRefreshCookie(w http.ResponseWriter, token, tokenName string, cfg Config, ttl time.Duration) {
	expires := time.Now().Add(ttl)

	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
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

func RemoveRefreshCookie(w http.ResponseWriter, tokenName string, cfg Config) {
	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
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
