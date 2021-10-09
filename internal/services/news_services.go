package services

import (
	"net/url"
	"stonks/internal/config"
	"stonks/internal/interfaces"
	"stonks/internal/models"
)

type NewsService struct {
	NewsRepo     interfaces.INewsRepo
	NewsHandler interfaces.INewsHandler
	Config *config.Config
}

func CreateUrl(ParameterMap url.Values) string {
	u := "https://api.newscatcherapi.com/v2/search?"
	mapLen := len(ParameterMap)
	i := 1
	for k, v := range ParameterMap {
		u += k + "=" + v[0]
		if i != mapLen {
			u += "&"
		}
		i += 1
	}
	return u
}

func (s *NewsService) GetNews(QueryMap url.Values) (models.News, error) {
	u := CreateUrl(QueryMap)
	client := s.NewsHandler.GetClient()
	cfg := s.NewsHandler.GetConfig()
	resp, err := s.NewsRepo.GetNewsR(cfg, client, u)
	if err != nil {
		return models.News{}, err
	}

	return resp, nil
}
