package infrastructure

import (
	"go.uber.org/zap"
	"stonks/internal/config"
	"stonks/internal/controllers"
	"stonks/internal/interfaces"
	"stonks/internal/repos"
	"stonks/internal/services"
)

type IInjector interface {
	InjectNewsController() controllers.NewsControllers
}

var env *environment

type environment struct {
	logger *zap.SugaredLogger
	httpClient interfaces.INewsHandler
	cfg *config.Config
}

func (e *environment) InjectNewsController() controllers.NewsControllers {
	return controllers.NewsControllers{
		Log: e.logger,
		NewsService: &services.NewsService{
			NewsRepo: &repos.NewsRepo{},
			NewsHandler: e.httpClient,
			Config: e.cfg,
		},

	}
}

func Injector(logger *zap.SugaredLogger, config *config.Config) (IInjector, error) {
	client := InitNewsAPIClient(config)
	env = &environment{
		logger: logger,
		httpClient: client,
		cfg: config,
	}
	return env, nil
}
