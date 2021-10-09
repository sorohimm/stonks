package interfaces

import (
	"net/http"
	"stonks/internal/config"
	"stonks/internal/models"
)

type INewsRepo interface {
	GetNewsR(*config.Config, http.Client, string) (models.News, error)
}
