package infrastructure

import (
	"net/http"
	"stonks/internal/config"
	"stonks/internal/interfaces"
)

type NewsClient struct {
	Client http.Client
	Config *config.Config
}

func InitNewsAPIClient(cfg *config.Config) interfaces.INewsHandler {
	client := http.Client{}
	return &NewsClient{Client: client, Config: cfg}
}

func (h *NewsClient) GetClient() http.Client {
	return h.Client
}

func (h *NewsClient) GetConfig() *config.Config {
	return h.Config
}