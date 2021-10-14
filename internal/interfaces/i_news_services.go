package interfaces

import (
	"net/url"
	"stonks/internal/models/news"
)

type INewsService interface {
	GetNews(url.Values) (news.News, error)
}

