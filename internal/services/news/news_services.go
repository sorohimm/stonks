package news

import (
	"net/http"
	"net/url"
	"stonks/internal/config"
	news2 "stonks/internal/models/news"

	"stonks/internal/constants/news"
	"stonks/internal/interfaces"
)

type NewsService struct {
	NewsRepo    interfaces.INewsRepo
	Config      *config.Config
}

func (s *NewsService) GetNews(queryParams url.Values) (news2.News, error) {
	req, _ := http.NewRequest(http.MethodGet, news.URL, nil)
	req.URL.Path = news.Path
	req.URL.RawQuery = queryParams.Encode()
	req.Header.Set("x-api-key", s.Config.NKey)

	resp, err := s.NewsRepo.GetNews(req)
	if err != nil {
		return news2.News{}, err
	}

	return resp, nil
}
