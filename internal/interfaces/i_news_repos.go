package interfaces

import (
	"net/http"
	"stonks/internal/models/news"
)

type INewsRepo interface {
	GetNews(*http.Request) (news.News, error)
}
