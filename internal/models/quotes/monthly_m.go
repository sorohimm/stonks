package quotes_models

type MonthlyTSMongo struct {
	Information   string `json:"information,omitempty"`
	Symbol        string `json:"symbol,omitempty"`
	LastRefreshed string `json:"last_refreshed,omitempty"`
	TimeZone      string `json:"time_zone,omitempty"`
	Series   []SessionMongo `json:"series,omitempty"`
}

func (d *MonthlyTSMongo) setMeta(v Meta) {
	d.Information = v.Information
	d.Symbol = v.Symbol
	d.LastRefreshed = v.LastRefreshed
	d.TimeZone = v.TimeZone
}

func (d *MonthlyTSMongo) Set(v MonthlyTS) {
	var mseries []SessionMongo
	for date, s := range v.Series {
		newSession := SessionMongo{
			Date:   date,
			Open:   s.Open,
			High:   s.High,
			Low:    s.Low,
			Close:  s.Close,
			Volume: s.Volume,
		}
		mseries = append(mseries, newSession)
	}
	d.setMeta(v.MetaData)
	d.Series = mseries
}
