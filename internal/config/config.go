package config

import "os"

type Config struct {
	StonksPort string
	NewsAuthData
	MarketAuthData
	DbAuthenticationData
}

type NewsAuthData struct {
	NewsKey string
}

type MarketAuthData struct {
	MarketKey string
}

type DbAuthenticationData struct {
	DbUsername string
	DbPassword string
	DbName     string
	DbHost     string
	DbPort     string
	DbURI      string
}

func New() *Config {
	return &Config{
		StonksPort: os.Getenv("STONKS_PORT"),
		NewsAuthData: NewsAuthData{
			NewsKey: os.Getenv("NEWS_KEY"),
		},
		MarketAuthData: MarketAuthData{
			MarketKey: os.Getenv("MARKET_KEY"),
		},
		DbAuthenticationData: DbAuthenticationData{
			DbUsername: os.Getenv("MONGODB_USER"),
			DbPassword: os.Getenv("MONGODB_PASSWORD"),
			DbHost:     os.Getenv("MONGODB_HOST"),
			DbPort:     os.Getenv("MONGODB_PORT"),
			DbName:     os.Getenv("MONGODB_NAME"),
			DbURI:      os.Getenv("MONGODB_URI"),
		},
	}
}
