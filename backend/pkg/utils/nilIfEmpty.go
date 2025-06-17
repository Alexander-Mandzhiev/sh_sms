package utils

func NilIfEmpty(s *string) interface{} {
	if s == nil || *s == "" {
		return nil
	}
	return *s
}
