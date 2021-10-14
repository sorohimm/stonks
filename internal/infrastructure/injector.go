package infrastructure

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"stonks/internal/config"
	ec "stonks/internal/controllers/market/earnings"
	isc "stonks/internal/controllers/market/income_statement"
	ovc "stonks/internal/controllers/market/overview"
	nc "stonks/internal/controllers/news"

	er "stonks/internal/repos/earnings"
	isr "stonks/internal/repos/income_statement"
	nr "stonks/internal/repos/news"
	or "stonks/internal/repos/overview"

	es "stonks/internal/services/earnings"
	iss "stonks/internal/services/income_statement"
	ns "stonks/internal/services/news"
	ovs "stonks/internal/services/overview"
)

type IInjector interface {
	InjectNewsController() nc.NewsControllers
	InjectOverviewController() ovc.OverviewControllers
	InjectEarningsController() ec.EarningsControllers
	InjectIncomeStatementController() isc.IncomeStatementControllers
}

var env *environment

type environment struct {
	logger *zap.SugaredLogger
	cfg    *config.Config
}

func (e *environment) InjectNewsController() nc.NewsControllers {
	return nc.NewsControllers{
		Log: e.logger,
		NewsService: &ns.NewsService{
			NewsRepo: &nr.NewsRepo{},
			Config:   e.cfg,
		},
		Validator: validator.New(),
	}
}

func (e *environment) InjectOverviewController() ovc.OverviewControllers {
	return ovc.OverviewControllers{
		Log: e.logger,
		OverviewService: &ovs.OverviewService{
			OverviewRepo: &or.OverviewRepo{},
			Config:       e.cfg,
		},
		Validator: validator.New(),
	}
}

func (e *environment) InjectEarningsController() ec.EarningsControllers {
	return ec.EarningsControllers{
		Log: e.logger,
		EarningsService: &es.EarningsService{
			EarningsRepo: &er.EarningsRepo{},
			Config:       e.cfg,
		},
		Validator: validator.New(),
	}
}

func (e *environment) InjectIncomeStatementController() isc.IncomeStatementControllers {
	return isc.IncomeStatementControllers{
		Log: e.logger,
		IncomeStatementService: &iss.IncomeStatementService{
			IncomeStatementRepo: &isr.IncomeStatementRepo{},
			Config:              e.cfg,
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
