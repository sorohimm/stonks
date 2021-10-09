package interfaces

import "stonks/internal/models"

type INewsService interface {
	GetNews(string) (models.News, error)
}

