package infrastructure

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"stonks/internal/controllers/market/stock"
	details_repo "stonks/internal/repos/details"
	details_service "stonks/internal/services/details"
	"stonks/internal/services/news"

	"stonks/internal/config"
	"stonks/internal/controllers/market/details"
	"stonks/internal/controllers/news"
	"stonks/internal/repos/news"
	"stonks/internal/repos/stock"
	"stonks/internal/services/stock"
)

type IInjector interface {
	InjectNewsController() news_controller.NewsControllers
	InjectDetailsController() details_controller.CompanyDetailsControllers
	InjectStockController() stock_controller.StockControllers
}

var env *environment

type environment struct {
	logger *zap.SugaredLogger
	cfg    *config.Config
}

func (e *environment) InjectDetailsController() details_controller.CompanyDetailsControllers {
	return details_controller.CompanyDetailsControllers{
		Log: e.logger,
		CompanyDetailsService: &details_service.CompanyDetailsService{
			DetailsRepo: &details_repo.CompanyDetailsRepo{},
			Config:      e.cfg,
		},
		Validator: validator.New(),
	}
}

func (e *environment) InjectNewsController() news_controller.NewsControllers {
	return news_controller.NewsControllers{
		Log: e.logger,
		NewsService: &news_service.NewsService{
			NewsRepo: &news_repo.NewsRepo{},
			Config:   e.cfg,
		},
		Validator: validator.New(),
	}
}

func (e *environment) InjectStockController() stock_controller.StockControllers {
	return stock_controller.StockControllers{
		Log: e.logger,
		StockService: &stock_service.StockService{
			StockRepo: &stock_repo.StockRepo{},
			Config:    e.cfg,
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
