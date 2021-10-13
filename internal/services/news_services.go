package services

import (
	"net/http"
	"net/url"
	"stonks/internal/config"
	"stonks/internal/constants/news"
	"stonks/internal/interfaces/news_interfaces"
	"stonks/internal/models/news"
)

type NewsService struct {
	NewsRepo news_interface.INewsRepo
	Config   *config.Config
}

func (s *NewsService) GetNews(queryParams url.Values) (news_models.News, error) {
	req, _ := http.NewRequest(http.MethodGet, news_constants.URL, nil)
	req.URL.Path = news_constants.Path
	req.URL.RawQuery = queryParams.Encode()
	req.Header.Set("x-api-key", s.Config.NKey)

	resp, err := s.NewsRepo.GetNews(req)
	if err != nil {
		return news_models.News{}, err
	}

	return resp, nil
}
