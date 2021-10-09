package interfaces

import (
	"net/http"
	"stonks/internal/config"
)

type INewsHandler interface {
	GetClient() http.Client
	GetConfig() *config.Config
}
