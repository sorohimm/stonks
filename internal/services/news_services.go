package services

import (
	"net/http"
	"net/url"
	"stonks/internal/config"

	"stonks/internal/constants/news"
	"stonks/internal/interfaces"
	"stonks/internal/models"
)

type NewsService struct {
	NewsRepo    interfaces.INewsRepo
	Config      *config.Config
}

func (s *NewsService) GetNews(queryParams url.Values) (models.News, error) {
	req, _ := http.NewRequest(http.MethodGet, news.URL, nil)
	req.URL.Path = news.Path
	req.URL.RawQuery = queryParams.Encode()
	req.Header.Set("x-api-key", s.Config.NKey)

	resp, err := s.NewsRepo.GetNews(req)
	if err != nil {
		return models.News{}, err
	}

	return resp, nil
}
