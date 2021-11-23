package details_models

import "strconv"

type Overview struct {
	Symbol                     string `json:"Symbol,omitempty"`
	AssetType                  string `json:"AssetType,omitempty"`
	Name                       string `json:"Name,omitempty"`
	Description                string `json:"Description,omitempty"`
	CIK                        string `json:"CIK,omitempty"`
	Exchange                   string `json:"Exchange,omitempty"`
	Currency                   string `json:"Currency,omitempty"`
	Country                    string `json:"Country,omitempty"`
	Sector                     string `json:"Sector,omitempty"`
	Industry                   string `json:"Industry,omitempty"`
	Address                    string `json:"Address,omitempty"`
	FiscalYearEnd              string `json:"FiscalYearEnd,omitempty"`
	LatestQuarter              string `json:"LatestQuarter,omitempty"`
	MarketCapitalization       string `json:"MarketCapitalization,omitempty"`
	EBITDA                     string `json:"EBITDA,omitempty"`
	PERatio                    string `json:"PERatio,omitempty"`
	PEGRatio                   string `json:"PEGRatio,omitempty"`
	BookValue                  string `json:"BookValue,omitempty"`
	DividendPerShare           string `json:"DividendPerShare,omitempty"`
	DividendYield              string `json:"DividendYield,omitempty"`
	EPS                        string `json:"EPS,omitempty"`
	RevenuePerShareTTM         string `json:"RevenuePerShareTTM,omitempty"`
	ProfitMargin               string `json:"ProfitMargin,omitempty"`
	OperatingMarginTTM         string `json:"OperatingMarginTTM,omitempty"`
	ReturnOnAssetsTTM          string `json:"ReturnOnAssetsTTM,omitempty"`
	ReturnOnEquityTTM          string `json:"ReturnOnEquityTTM,omitempty"`
	RevenueTTM                 string `json:"RevenueTTM,omitempty"`
	GrossProfitTTM             string `json:"GrossProfitTTM,omitempty"`
	DilutedEPSTTM              string `json:"DilutedEPSTTM,omitempty"`
	QuarterlyEarningsGrowthYOY string `json:"QuarterlyEarningsGrowthYOY,omitempty"`
	QuarterlyRevenueGrowthYOY  string `json:"QuarterlyRevenueGrowthYOY,omitempty"`
	AnalystTargetPrice         string `json:"AnalystTargetPrice,omitempty"`
	TrailingPE                 string `json:"TrailingPE,omitempty"`
	ForwardPE                  string `json:"ForwardPE,omitempty"`
	PriceToSalesRatioTTM       string `json:"PriceToSalesRatioTTM,omitempty"`
	PriceToBookRatio           string `json:"PriceToBookRatio,omitempty"`
	EVToRevenue                string `json:"EVToRevenue,omitempty"`
	EVToEBITDA                 string `json:"EVToEBITDA,omitempty"`
	Beta                       string `json:"Beta,omitempty"`
	WeekHigh                   string `json:"52WeekHigh,omitempty"`
	WeekLow                    string `json:"52WeekLow,omitempty"`
	DayMovingAverage           string `json:"50DayMovingAverage,omitempty"`
	DayMovingAverage1          string `json:"200DayMovingAverage,omitempty"`
	SharesOutstanding          string `json:"SharesOutstanding,omitempty"`
	DividendDate               string `json:"DividendDate,omitempty"`
	ExDividendDate             string `json:"ExDividendDate,omitempty"`
}

type OverviewMongo struct {
	Symbol                     string  `json:"symbol,omitempty"`
	AssetType                  string  `json:"AssetType,omitempty"`
	Name                       string  `json:"Name,omitempty"`
	Description                string  `json:"Description,omitempty"`
	CIK                        string  `json:"CIK,omitempty"`
	Exchange                   string  `json:"Exchange,omitempty"`
	Currency                   string  `json:"Currency,omitempty"`
	Country                    string  `json:"Country,omitempty"`
	Sector                     string  `json:"Sector,omitempty"`
	Industry                   string  `json:"Industry,omitempty"`
	Address                    string  `json:"Address,omitempty"`
	FiscalYearEnd              string  `json:"FiscalYearEnd,omitempty"`
	LatestQuarter              string  `json:"LatestQuarter,omitempty"`
	MarketCapitalization       float64 `json:"MarketCapitalization,omitempty"`
	EBITDA                     float64 `json:"EBITDA,omitempty"`
	PERatio                    float64 `json:"PERatio,omitempty"`
	PEGRatio                   float64 `json:"PEGRatio,omitempty"`
	BookValue                  float64 `json:"BookValue,omitempty"`
	DividendPerShare           float64 `json:"DividendPerShare,omitempty"`
	DividendYield              float64 `json:"DividendYield,omitempty"`
	EPS                        float64 `json:"EPS,omitempty"`
	RevenuePerShareTTM         float64 `json:"RevenuePerShareTTM,omitempty"`
	ProfitMargin               float64 `json:"ProfitMargin,omitempty"`
	OperatingMarginTTM         float64 `json:"OperatingMarginTTM,omitempty"`
	ReturnOnAssetsTTM          float64 `json:"ReturnOnAssetsTTM,omitempty"`
	ReturnOnEquityTTM          float64 `json:"ReturnOnEquityTTM,omitempty"`
	RevenueTTM                 float64 `json:"RevenueTTM,omitempty"`
	GrossProfitTTM             float64 `json:"GrossProfitTTM,omitempty"`
	DilutedEPSTTM              float64 `json:"DilutedEPSTTM,omitempty"`
	QuarterlyEarningsGrowthYOY float64 `json:"QuarterlyEarningsGrowthYOY,omitempty"`
	QuarterlyRevenueGrowthYOY  float64 `json:"QuarterlyRevenueGrowthYOY,omitempty"`
	AnalystTargetPrice         float64 `json:"AnalystTargetPrice,omitempty"`
	TrailingPE                 float64 `json:"TrailingPE,omitempty"`
	ForwardPE                  float64 `json:"ForwardPE,omitempty"`
	PriceToSalesRatioTTM       float64 `json:"PriceToSalesRatioTTM,omitempty"`
	PriceToBookRatio           float64 `json:"PriceToBookRatio,omitempty"`
	EVToRevenue                float64 `json:"EVToRevenue,omitempty"`
	EVToEBITDA                 float64 `json:"EVToEBITDA,omitempty"`
	Beta                       float64 `json:"Beta,omitempty"`
	WeekHigh                   float64 `json:"52WeekHigh,omitempty"`
	WeekLow                    float64 `json:"52WeekLow,omitempty"`
	DayMovingAverage           float64 `json:"50DayMovingAverage,omitempty"`
	DayMovingAverage1          float64 `json:"200DayMovingAverage,omitempty"`
	SharesOutstanding          float64 `json:"SharesOutstanding,omitempty"`
	DividendDate               string  `json:"DividendDate,omitempty"`
	ExDividendDate             string  `json:"ExDividendDate,omitempty"`
}

func (m *OverviewMongo) Set(v Overview) {
	MarketCapitalization, _ := strconv.ParseFloat(v.MarketCapitalization, 64)
	EBITDA, _ := strconv.ParseFloat(v.EBITDA, 64)
	PERatio, _ := strconv.ParseFloat(v.PERatio, 64)
	PEGRatio, _ := strconv.ParseFloat(v.PEGRatio, 64)
	BookValue, _ := strconv.ParseFloat(v.BookValue, 64)
	DividendPerShare, _ := strconv.ParseFloat(v.DividendPerShare, 64)
	DividendYield, _ := strconv.ParseFloat(v.DividendYield, 64)
	EPS, _ := strconv.ParseFloat(v.EPS, 64)
	RevenuePerShareTTM, _ := strconv.ParseFloat(v.RevenuePerShareTTM, 64)
	ProfitMargin, _ := strconv.ParseFloat(v.ProfitMargin, 64)
	OperatingMarginTTM, _ := strconv.ParseFloat(v.OperatingMarginTTM, 64)
	ReturnOnAssetsTTM, _ := strconv.ParseFloat(v.ReturnOnAssetsTTM, 64)
	ReturnOnEquityTTM, _ := strconv.ParseFloat(v.ReturnOnEquityTTM, 64)
	RevenueTTM, _ := strconv.ParseFloat(v.RevenueTTM, 64)
	GrossProfitTTM, _ := strconv.ParseFloat(v.GrossProfitTTM, 64)
	DilutedEPSTTM, _ := strconv.ParseFloat(v.DilutedEPSTTM, 64)
	QuarterlyEarningsGrowthYOY, _ := strconv.ParseFloat(v.QuarterlyEarningsGrowthYOY, 64)
	QuarterlyRevenueGrowthYOY, _ := strconv.ParseFloat(v.QuarterlyRevenueGrowthYOY, 64)
	AnalystTargetPrice, _ := strconv.ParseFloat(v.AnalystTargetPrice, 64)
	TrailingPE, _ := strconv.ParseFloat(v.TrailingPE, 64)
	ForwardPE, _ := strconv.ParseFloat(v.ForwardPE, 64)
	PriceToSalesRatioTTM, _ := strconv.ParseFloat(v.PriceToSalesRatioTTM, 64)
	PriceToBookRatio, _ := strconv.ParseFloat(v.PriceToBookRatio, 64)
	EVToRevenue, _ := strconv.ParseFloat(v.EVToRevenue, 64)
	EVToEBITDA, _ := strconv.ParseFloat(v.EVToEBITDA, 64)
	Beta, _ := strconv.ParseFloat(v.Beta, 64)
	WeekHigh, _ := strconv.ParseFloat(v.WeekHigh, 64)
	WeekLow, _ := strconv.ParseFloat(v.WeekLow, 64)
	DayMovingAverage, _ := strconv.ParseFloat(v.DayMovingAverage, 64)
	DayMovingAverage1, _ := strconv.ParseFloat(v.DayMovingAverage1, 64)
	SharesOutstanding, _ := strconv.ParseFloat(v.SharesOutstanding, 64)

	m.Symbol = v.Symbol
	m.AssetType = v.AssetType
	m.Name = v.Name
	m.Description = v.Description
	m.CIK = v.CIK
	m.Exchange = v.Exchange
	m.Currency = v.Currency
	m.Country = v.Country
	m.Sector = v.Sector
	m.Industry = v.Industry
	m.Address = v.Address
	m.FiscalYearEnd = v.FiscalYearEnd
	m.LatestQuarter = v.LatestQuarter
	m.MarketCapitalization = MarketCapitalization
	m.EBITDA = EBITDA
	m.PERatio = PERatio
	m.PEGRatio = PEGRatio
	m.BookValue = BookValue
	m.DividendPerShare = DividendPerShare
	m.DividendYield = DividendYield
	m.EPS = EPS
	m.RevenuePerShareTTM = RevenuePerShareTTM
	m.ProfitMargin = ProfitMargin
	m.OperatingMarginTTM = OperatingMarginTTM
	m.ReturnOnAssetsTTM = ReturnOnAssetsTTM
	m.ReturnOnEquityTTM = ReturnOnEquityTTM
	m.RevenueTTM = RevenueTTM
	m.GrossProfitTTM = GrossProfitTTM
	m.DilutedEPSTTM = DilutedEPSTTM
	m.QuarterlyEarningsGrowthYOY = QuarterlyEarningsGrowthYOY
	m.QuarterlyRevenueGrowthYOY = QuarterlyRevenueGrowthYOY
	m.AnalystTargetPrice = AnalystTargetPrice
	m.TrailingPE = TrailingPE
	m.ForwardPE = ForwardPE
	m.PriceToSalesRatioTTM = PriceToSalesRatioTTM
	m.PriceToBookRatio = PriceToBookRatio
	m.EVToRevenue = EVToRevenue
	m.EVToEBITDA = EVToEBITDA
	m.Beta = Beta
	m.WeekHigh = WeekHigh
	m.WeekLow = WeekLow
	m.DayMovingAverage = DayMovingAverage
	m.DayMovingAverage1 = DayMovingAverage1
	m.SharesOutstanding = SharesOutstanding
	m.DividendDate = v.DividendDate
	m.ExDividendDate = v.ExDividendDate
}
