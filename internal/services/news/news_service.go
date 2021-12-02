package news_service

import (
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"stonks/internal/config"
	"stonks/internal/models/news"

	"stonks/internal/constants/news"
	"stonks/internal/interfaces/news_interfaces"
)

type NewsService struct {
	Log      *zap.SugaredLogger
	NewsRepo news_interface.INewsRepo
	Config   *config.Config
}

func (s *NewsService) GetNews(queryParams url.Values) (news_models.News, error) {
	req, _ := http.NewRequest(http.MethodGet, news_constants.URL, nil)
	req.URL.Path = news_constants.Path
	req.URL.RawQuery = queryParams.Encode()
	req.Header.Set("x-stocks_api-key", s.Config.NewsKey)

	resp, err := s.NewsRepo.GetNews(req)
	if err != nil {
		return news_models.News{}, err
	}

	return resp, nil
}
