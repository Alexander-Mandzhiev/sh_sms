package service

func maskSecret(secret string) string {
	if len(secret) < 4 {
		return "****"
	}
	return secret[:2] + "****" + secret[len(secret)-2:]
}
