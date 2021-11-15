package quotes_models

type IntradayTSMongo struct {
	Information   string `json:"information,omitempty"`
	Symbol        string `json:"symbol,omitempty"`
	LastRefreshed string `json:"last_refreshed,omitempty" `
	TimeZone      string `json:"time_zone,omitempty" `
	Series   []SessionMongo `json:"series,omitempty" bson:"series,omitempty"`
}

func (i *IntradayTSMongo) setMeta(v Meta) {
	i.Information = v.Information
	i.Symbol = v.Symbol
	i.LastRefreshed = v.LastRefreshed
	i.TimeZone = v.TimeZone
}

func (i *IntradayTSMongo) Set1(v Intraday1TS) {
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
	i.setMeta(v.MetaData)
	i.Series = mseries
}

func (i *IntradayTSMongo) Set5(v Intraday5TS) {
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
	i.setMeta(v.MetaData)
	i.Series = mseries
}

func (i *IntradayTSMongo) Set15(v Intraday15TS) {
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
	i.setMeta(v.MetaData)
	i.Series = mseries
}

func (i *IntradayTSMongo) Set30(v Intraday30TS) {
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
	i.setMeta(v.MetaData)
	i.Series = mseries
}

func (i *IntradayTSMongo) Set60(v Intraday60TS) {
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
	i.setMeta(v.MetaData)
	i.Series = mseries
}
