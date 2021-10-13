package news_interface

import (
	"net/http"
	"stonks/internal/models/news"
)

type INewsRepo interface {
	GetNews(*http.Request) (news_models.News, error)
}
