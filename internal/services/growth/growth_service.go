package growth_services

import (
	"fmt"
	"go.uber.org/zap"
	"net/url"
	"stonks/internal/config"
	"stonks/internal/db/filter"
	"stonks/internal/interfaces/db_interfaces"
	"stonks/internal/interfaces/growth_interfaces"
	"stonks/internal/models"
	gmodels "stonks/internal/models/growth"
	"strconv"
)

type GrowthService struct {
	Log    *zap.SugaredLogger
	Config *config.Config
	GrowthRepo growth_interfaces.IGrowthRepo
	DbHandler db_interfaces.IDBHandler
}

func growthPercent(begin string, end string) string {
	bf, _ := strconv.ParseFloat(begin, 32)
	ef, _ := strconv.ParseFloat(end, 32)
	percent := (ef / bf) * 100 - 100
	return fmt.Sprintf("%.3f%s", percent, " %")
}

func CalculateGrowth(v *gmodels.Response) {
	v.Growth.OpenGrowth = growthPercent(v.Begin.Open, v.End.Open)
	v.Growth.HighGrowth = growthPercent(v.Begin.High, v.End.High)
	v.Growth.LowGrowth = growthPercent(v.Begin.Low, v.End.Low)
	v.Growth.CloseGrowth = growthPercent(v.Begin.Close, v.End.Close)
}

func (g *GrowthService) GetGrowth(values url.Values) (interface{}, error) {
	database := g.DbHandler.AcquireDatabase(g.Config.DbName)
	var t models.Timing
	t.Set(values)
	var pipe = filter.GrowthPipeline(t)
	response, err := g.GrowthRepo.GetQuote(database, "DailySeries", pipe)
	if err != nil {
		g.Log.Infof("growth_service :: dbroutine error")
		return nil, err
	}

	CalculateGrowth(&response)

	return response, nil
}
