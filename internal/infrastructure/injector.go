package infrastructure

import (
	"context"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"log"
	"stonks/internal/controllers/market/quotes"
	"stonks/internal/interfaces/db_interfaces"
	 "stonks/internal/repos/details"
	 "stonks/internal/services/details"
	"stonks/internal/services/news"

	"stonks/internal/config"
	"stonks/internal/controllers/market/details"
	"stonks/internal/controllers/news"
	"stonks/internal/repos/news"
	"stonks/internal/repos/quotes"
	"stonks/internal/services/quotes"
)

type IInjector interface {
	InjectNewsController() news_controller.NewsControllers
	InjectDetailsController() details_controller.CompanyDetailsControllers
	InjectQuotesController() quotes_controller.QuotesControllers
}

var env *environment

type environment struct {
	logger   *zap.SugaredLogger
	cfg      *config.Config
	dbClient db_interfaces.IDBHandler
}

func (e *environment) InjectDetailsController() details_controller.CompanyDetailsControllers {
	log.Println("inject details done")
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
	log.Println("inject news done")
	return news_controller.NewsControllers{
		Log: e.logger,
		NewsService: &news_service.NewsService{
			NewsRepo: &news_repo.NewsRepo{},
			Config:   e.cfg,
		},
		Validator: validator.New(),
	}
}

func (e *environment) InjectQuotesController() quotes_controller.QuotesControllers {
	log.Println("inject quotes done")
	return quotes_controller.QuotesControllers{
		Log: e.logger,
		QuotesService: &quotes_service.QuotesService{
			StockRepo: &quotes_repo.QuotesRepo{},
			Config:    e.cfg,
		},
		Validator: validator.New(),
	}
}

func Injector(logger *zap.SugaredLogger, ctx context.Context, cfg *config.Config) (IInjector, error) {
	logger.Infof("injector starting...")
	dbClient, err := InitDbClient(logger, cfg, ctx)
	if err != nil {
		logger.Infof("db init error")
		return nil, err
	}
	logger.Infof("db init ok")

	env = &environment{
		logger:   logger,
		cfg:      cfg,
		dbClient: dbClient,
	}

	log.Println("injecting done")
	return env, nil
}
