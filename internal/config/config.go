package config

import "os"

type Config struct {
	Port string

	NAPIAuthData
}

type NAPIAuthData struct {
	Key string
}

func New() (*Config, error) {
	return &Config{
		Port: os.Getenv("PORT"),
		NAPIAuthData: NAPIAuthData{
			Key: os.Getenv("NCAPI_KEY"),
		},
	}, nil
}
