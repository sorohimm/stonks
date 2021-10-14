package news_interface

import (
	"net/http"
	nm "stonks/internal/models/news"
)

type INewsRepo interface {
	GetNews(*http.Request) (nm.News, error)
}
