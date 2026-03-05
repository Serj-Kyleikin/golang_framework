package config

import (
	"os"
)

func Env(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
