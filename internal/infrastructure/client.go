package infrastructure

import (
	"stonks/internal/config"
	"stonks/internal/interfaces/client_interfaces"
)

type StonksConfig struct {
	Config *config.Config
}

func InitNewsAPIClient(cfg *config.Config) client_interfaces.INewsHandler {
	return &StonksConfig{Config: cfg}
}

func (h *StonksConfig) GetConfig() *config.Config {
	return h.Config
}
