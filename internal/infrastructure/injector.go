package infrastructure

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"stonks/internal/config"
	"stonks/internal/controllers"
	"stonks/internal/repos"
	"stonks/internal/services"
)

type IInjector interface {
	InjectNewsController() controllers.NewsControllers
	InjectOverviewController() controllers.OverviewControllers
	InjectEarningsController() controllers.EarningsControllers
}

var env *environment

type environment struct {
	logger *zap.SugaredLogger
	cfg    *config.Config
}

func (e *environment) InjectNewsController() controllers.NewsControllers {
	return controllers.NewsControllers{
		Log: e.logger,
		NewsService: &services.NewsService{
			NewsRepo: &repos.NewsRepo{},
			Config:   e.cfg,
		},
		Validator: validator.New(),
	}
}

func (e *environment) InjectOverviewController() controllers.OverviewControllers {
	return controllers.OverviewControllers{
		Log: e.logger,
		OverviewService: &services.OverviewService{
			OverviewRepo: &repos.OverviewRepo{},
			Config:       e.cfg,
		},
		Validator: validator.New(),
	}
}
func (e *environment) InjectEarningsController() controllers.EarningsControllers {
	return controllers.EarningsControllers{
		Log: e.logger,
		EarningsService: &services.EarningsService{
			EarningsRepo: &repos.EarningsRepo{},
			Config:       e.cfg,
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
