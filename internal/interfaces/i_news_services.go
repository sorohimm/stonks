package interfaces

import (
	"net/url"
	"stonks/internal/models"
)

type INewsService interface {
	GetNews(url.Values) (models.News, error)

}

