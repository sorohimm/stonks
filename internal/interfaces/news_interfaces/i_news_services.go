package news_interface

import (
	"net/url"
	nm "stonks/internal/models/news"
)

type INewsService interface {
	GetNews(url.Values) (nm.News, error)
}

