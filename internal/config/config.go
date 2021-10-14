package config

import "os"

type Config struct {
	Port string
	NewsAuthData
	MarketAuthData
}

type NewsAuthData struct {
	NewsKey string
}

type MarketAuthData struct {
	MarketKey string
}

func New() *Config {
	return &Config{
		Port: getEnv("PORT", ""),
		NewsAuthData: NewsAuthData{
			NewsKey: getEnv("NEWS_KEY", ""),
		},
		MarketAuthData: MarketAuthData{
			MarketKey: getEnv("MARKET_KEY", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
