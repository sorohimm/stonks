package services

import (
	"net/http"
	"net/url"
	"stonks/internal/config"
	"stonks/internal/interfaces"
	"stonks/internal/models"
)

type NewsService struct {
	NewsRepo    interfaces.INewsRepo
	NewsHandler interfaces.INewsHandler
	Config      *config.Config
}

func (s *NewsService) GetNews(queryParams url.Values) (models.News, error) {
	u := "https://api.newscatcherapi.com"
	p := "/v2/search"

	req, _ := http.NewRequest(http.MethodGet, u, nil)
	req.URL.Path = p
	req.URL.RawQuery = queryParams.Encode()
	req.Header.Set("x-api-key", s.Config.NKey)

	resp, err := s.NewsRepo.GetNews(req)
	if err != nil {
		return models.News{}, err
	}

	return resp, nil
}
