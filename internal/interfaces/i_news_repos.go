package interfaces

import (
	"net/http"
	"stonks/internal/models"
)

type INewsRepo interface {
	GetNews(*http.Request) (models.News, error)
}
