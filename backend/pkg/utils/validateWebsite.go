package utils

import (
	"fmt"
	"net/url"
)

func ValidateWebsite(website string, length int) error {
	if err := ValidateString(website, length); err != nil {
		return err
	}

	u, err := url.Parse(website)
	if err != nil {
		return fmt.Errorf("invalid URL format")
	}
	if u.Scheme == "" || u.Host == "" {
		return fmt.Errorf("URL must contain scheme and host")
	}
	return nil
}
