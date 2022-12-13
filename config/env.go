package config

import "os"

const (
	ENV = "ENV"
)

func IsLocal() bool {
	env := os.Getenv(ENV)
	return env == "" || env == "local"
}
