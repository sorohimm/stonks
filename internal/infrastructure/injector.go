package infrastructure

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"stonks/internal/config"
	"stonks/internal/controllers/news"
	"stonks/internal/repos"
	news2 "stonks/internal/services/news"
)

type IInjector interface {
	InjectNewsController() news.NewsControllers
}

var env *environment

type environment struct {
	logger *zap.SugaredLogger
	cfg    *config.Config
}

func (e *environment) InjectNewsController() news.NewsControllers {
	return news.NewsControllers{
		Log: e.logger,
		NewsService: &news2.NewsService{
			NewsRepo: &repos.NewsRepo{},
			Config:   e.cfg,
		},
		Validator: validator.New(),
	}
}

func Injector(logger *zap.SugaredLogger, config *config.Config) (IInjector, error) {
	env = &environment{
		logger: logger,
		cfg:    config,
	}
	return env, nil
}
