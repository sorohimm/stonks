package services

import (
	"log"
	"stonks/internal/interfaces"
	"stonks/internal/models"
)

type NewsService struct {
	NewsRepo     interfaces.INewsRepo
	//NewsKHandler interfaces.INewsRepo
}

func (s *NewsService) GetNews(url string) (models.News, error) {
	//TODO: reque valid
	resp, err := s.NewsRepo.GetNewsR(url)
	if err != nil {
		log.Print("request error", err.Error())
		return models.News{}, err
	}

	return resp, nil
}
