package details_models

type Overview struct {
	Symbol               string `json:"Symbol,omitempty"`
	AssetType            string `json:"AssetType,omitempty"`
	Name                 string `json:"Name,omitempty"`
	Description          string `json:"Description,omitempty"`
	Exchange             string `json:"Exchange,omitempty"`
	Currency             string `json:"Currency,omitempty"`
	Country              string `json:"Country,omitempty"`
	Sector               string `json:"Sector,omitempty"`
	Industry             string `json:"Industry,omitempty"`
	FiscalYearEnd        string `json:"FiscalYearEnd,omitempty"`
	DividendPerShare     string `json:"DividendPerShare,omitempty"`
	ProfitMargin         string `json:"ProfitMargin,omitempty"`
	MarketCapitalization string `json:"MarketCapitalization,omitempty"`
	BookValue            string `json:"BookValue,omitempty"`
	WeekHigh52           string `json:"52WeekHigh,omitempty"`
	WeekLow52            string `json:"52WeeLow,omitempty"`
	DayMovingAverage50   string `json:"50DayMovingAverage,omitempty"`
	DayMovingAverage200  string `json:"200DayMovingAverage,omitempty"`
	PayoutRatio          string `json:"PayoutRatio,omitempty"`
	DividendDate         string `json:"DividendDate,omitempty"`
}
