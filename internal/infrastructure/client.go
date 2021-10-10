package infrastructure

import (
	"stonks/internal/config"
	"stonks/internal/interfaces"
)

type NewsClient struct {
	Config *config.Config
}

func InitNewsAPIClient(cfg *config.Config) interfaces.INewsHandler {
	return &NewsClient{Config: cfg}
}

func (h *NewsClient) GetConfig() *config.Config {
	return h.Config
}
