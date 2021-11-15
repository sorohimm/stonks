package quotes_models

type WeeklyTSMongo struct {
	Information   string `json:"information,omitempty"`
	Symbol        string `json:"symbol,omitempty"`
	LastRefreshed string `json:"last_refreshed,omitempty" `
	TimeZone      string `json:"time_zone,omitempty" `
	Series   []SessionMongo `json:"series,omitempty"`
}

func (d *WeeklyTSMongo) setMeta(v Meta) {
	d.Information = v.Information
	d.Symbol = v.Symbol
	d.LastRefreshed = v.LastRefreshed
	d.TimeZone = v.TimeZone
}

func (d *WeeklyTSMongo) Set(v WeeklyTS) {
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
