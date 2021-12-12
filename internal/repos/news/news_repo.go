package news_repo

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"

	nm "stonks/internal/models/news"
)

type NewsRepo struct {
	Log    *zap.SugaredLogger
	Client *http.Client
}

func (r *NewsRepo) GetNews(request *http.Request) (nm.News, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != http.StatusOK {
		r.Log.Info("Request error")
		return nm.News{}, err
	}

	newsBody := nm.News{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&newsBody)
	if err != nil {
		r.Log.Info("Decode error")
		return nm.News{}, err
	}
	return newsBody, nil
}
