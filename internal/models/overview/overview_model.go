package overview_models

type Overview struct {
	Symbol               string `json:"Symbol"`
	AssetType            string `json:"AssetType"`
	Name                 string `json:"Name"`
	Description          string `json:"Description"`
	Exchange             string `json:"Exchange"`
	Currency             string `json:"Currency"`
	Country              string `json:"Country"`
	Sector               string `json:"Sector"`
	Industry             string `json:"Industry"`
	FiscalYearEnd        string `json:"FiscalYearEnd"`
	DividendPerShare     string `json:"DividendPerShare"`
	ProfitMargin         string `json:"ProfitMargin"`
	MarketCapitalization string `json:"MarketCapitalization"`
	BookValue            string `json:"BookValue"`
	WeekHigh52           string `json:"52WeekHigh"`
	WeekLow52            string `json:"52WeeLow"`
	DayMovingAverage50   string `json:"50DayMovingAverage"`
	DayMovingAverage200  string `json:"200DayMovingAverage"`
	PayoutRatio          string `json:"PayoutRatio"`
	DividendDate         string `json:"DividendDate"`
}

type Request struct {
	Symbol string `validate:"required"`
}
