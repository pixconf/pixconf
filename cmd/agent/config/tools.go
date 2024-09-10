package config

import "os"

func getEnvOrDefault(key, def string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return def
}

func getEnvOrDefaultFunc(key string, def func() string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return def()
}
