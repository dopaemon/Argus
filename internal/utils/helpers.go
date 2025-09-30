package utils

func GetString(fn func() (string, error)) string {
	if val, err := fn(); err == nil {
		return val
	}
	return ""
}
