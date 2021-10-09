package config

import "os"

type Config struct {
	Port string
	NAPIAuthData
}

type NAPIAuthData struct {
	NKey string
}

func New() *Config {
	return &Config{
		Port: getEnv("PORT", ""),
		NAPIAuthData: NAPIAuthData{
			NKey: getEnv("NCAPI_KEY", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
