package infrastructure

import (
	"go.uber.org/zap"
	"stonks/internal/config"
	"stonks/internal/controllers"
	"stonks/internal/repos"
	"stonks/internal/services"
)

type IInjector interface {
	InjectNewsController() controllers.NewsControllers
}

var env *environment

type environment struct {
	logger *zap.SugaredLogger
	config *config.Config
}

func (e *environment) InjectNewsController() controllers.NewsControllers {
	return controllers.NewsControllers{
		Log: e.logger,
		NewsService: &services.NewsService{
			NewsRepo: &repos.NewsRepo{},
		},
		Config: e.config,
	}
}

func Injector(logger *zap.SugaredLogger, config *config.Config) (IInjector, error) {
	env = &environment{
		logger: logger,
		config: config,
	}
	return env, nil
}
