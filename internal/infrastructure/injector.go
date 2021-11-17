package infrastructure

import (
	"context"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	growth_controller "stonks/internal/controllers/market/growth"
	"stonks/internal/controllers/market/quotes"
	"stonks/internal/interfaces/db_interfaces"
	"stonks/internal/repos/details"
	growth_repo "stonks/internal/repos/growth"
	"stonks/internal/services/details"
	growth_services "stonks/internal/services/growth"
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
	InjectGrowthController() growth_controller.GrowthControllers
}

var env *environment

type environment struct {
	logger   *zap.SugaredLogger
	cfg      *config.Config
	client *http.Client
	dbClient db_interfaces.IDBHandler
}

func (e *environment) InjectNewsController() news_controller.NewsControllers {
	e.logger.Info("inject news...")
	return news_controller.NewsControllers{
		Log: e.logger,
		NewsService: &news_service.NewsService{
			Log: e.logger,
			NewsRepo: &news_repo.NewsRepo{
				Client: http.DefaultClient,
			},
			Config:   e.cfg,
		},
		Validator: validator.New(),
	}
}

func (e *environment) InjectDetailsController() details_controller.CompanyDetailsControllers {
	e.logger.Info("inject details...")
	return details_controller.CompanyDetailsControllers{
		Log: e.logger,
		CompanyDetailsService: &details_service.CompanyDetailsService{
			Log: e.logger,
			Config:      e.cfg,
			DetailsRepo: &details_repo.CompanyDetailsRepo{
				Log: e.logger,
				Client: http.DefaultClient,
			},
			DbHandler: e.dbClient,
		},
		Validator: validator.New(),
	}
}

func (e *environment) InjectQuotesController() quotes_controller.QuotesControllers {
	e.logger.Info("inject quotes...")
	return quotes_controller.QuotesControllers{
		Log: e.logger,
		QuotesService: &quotes_service.QuotesService{
			Log: e.logger,
			Config:    e.cfg,
			QuotesRepo: &quotes_repo.QuotesRepo{
				Log: e.logger,
				Client: http.DefaultClient,
			},
			DbHandler: e.dbClient,
		},
		Validator: validator.New(),
	}
}

func (e *environment) InjectGrowthController() growth_controller.GrowthControllers {
	e.logger.Info("inject growth...")
	return growth_controller.GrowthControllers{
		Log: e.logger,
		GrowthService: &growth_services.GrowthService{
			Log: e.logger,
			Config:    e.cfg,
			GrowthRepo: &growth_repo.GrowthRepo{
				Log: e.logger,
				Client: http.DefaultClient,
			},
			DbHandler: e.dbClient,
		},
		Validator: validator.New(),
	}
}

func Injector(log *zap.SugaredLogger, ctx context.Context, cfg *config.Config) (IInjector, error) {
	log.Infof("injector starting...")
	dbClient, err := InitDbClient(log, cfg, ctx)
	if err != nil {
		log.Fatal("db init error")
		return nil, err
	}
	log.Infof("db init ok")

	env = &environment{
		logger:   log,
		cfg:      cfg,
		client:   http.DefaultClient,
		dbClient: dbClient,
	}

	log.Info("injecting done")
	return env, nil
}
