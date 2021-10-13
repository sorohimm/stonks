package news_interface

import (
	"net/url"
	"stonks/internal/models/news"
)

type INewsService interface {
	GetNews(url.Values) (news_models.News, error)
}

