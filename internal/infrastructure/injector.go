package infrastructure

import (
	"context"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	choose_controller "stonks/internal/controllers/market/choose"
	"stonks/internal/controllers/market/growth"
	"stonks/internal/controllers/market/quotes"
	"stonks/internal/interfaces/db_interfaces"
	"stonks/internal/repos/api"
	choose_repo "stonks/internal/repos/choose"
	db_repo "stonks/internal/repos/db"
	"stonks/internal/repos/details"
	"stonks/internal/repos/growth"

	"stonks/internal/services/choose"
	"stonks/internal/services/details"
	"stonks/internal/services/growth"
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
	InjectChooseController() choose_controller.ChooseControllers
}

var env *environment

type environment struct {
	logger   *zap.SugaredLogger
	cfg      *config.Config
	client *http.Client
	dbClient db_interfaces.IDBHandler
}

func (e *environment) InjectNewsController() news_controller.NewsControllers {
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
			QuotesApiRepo: &api_repo.ApiRepo{
				Log: e.logger,
				Client: http.DefaultClient,
			},
			DbRepo: &db_repo.DbRepo{
				Log: e.logger,
				Client: e.client,
			},
		},
		Validator: validator.New(),
	}
}

func (e *environment) InjectGrowthController() growth_controller.GrowthControllers {
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
			QuotesApiRepo: &api_repo.ApiRepo{
				Log: e.logger,
				Client: http.DefaultClient,
			},
			DbRepo: &db_repo.DbRepo{
				Log: e.logger,
				Client: e.client,
			},
		},
		Validator: validator.New(),
	}
}

func (e *environment) InjectChooseController() choose_controller.ChooseControllers {
	return choose_controller.ChooseControllers{
		Log: e.logger,
		ChooseService: &choose_service.ChooseService{
			Log: e.logger,
			Config:    e.cfg,
			ChooseRepo: &choose_repo.ChooseRepo{
				Log: e.logger,
				Client: http.DefaultClient,
			},
			DbHandler: e.dbClient,
			QuotesApiRepo: &api_repo.ApiRepo{
				Log: e.logger,
				Client: http.DefaultClient,
			},
			DbRepo: &db_repo.DbRepo{
				Log: e.logger,
				Client: e.client,
			},
		},
		Validator: validator.New(),
	}
}

func Injector(log *zap.SugaredLogger, ctx context.Context, cfg *config.Config) (IInjector, error) {
	log.Infof("injector starting...")
	client, err := InitDbClient(log, cfg, ctx)
	if err != nil {
		log.Fatal("db init error")
		return nil, err
	}
	log.Infof("db init ok")

	env = &environment{
		logger:   log,
		cfg:      cfg,
		client:   http.DefaultClient,
		dbClient: client,
	}

	log.Info("injecting done")
	return env, nil
}
