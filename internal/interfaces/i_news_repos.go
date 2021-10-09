package interfaces

import "stonks/internal/models"

type INewsRepo interface {
	GetNewsR(string) (models.News, error)
}
