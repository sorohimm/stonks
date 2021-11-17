package quotes_models

import (
	"sort"
)

type TSMongo struct {
	MetaData MetaMongo      `json:"_meta,omitempty" bson:"_meta,omitempty"`
	Series   []SessionMongo `json:"series,omitempty"`
}

func (d *TSMongo) sort() {
	sort.Slice(d.Series, func(i, j int) bool { return d.Series[i].Date > d.Series[j].Date })
}

func (d *TSMongo) SetD(v DailyTS) {
	var msessions []SessionMongo
	for date, s := range v.Series {
		newSession := SessionMongo{
			Date:   date,
			Open:   s.Open,
			High:   s.High,
			Low:    s.Low,
			Close:  s.Close,
			Volume: s.Volume,
		}
		msessions = append(msessions, newSession)
	}

	d.MetaData = MetaMongo(v.MetaData)
	d.Series = msessions
	d.sort()
}

func (d *TSMongo) SetW(v WeeklyTS) {
	var msessions []SessionMongo
	for date, s := range v.Series {
		newSession := SessionMongo{
			Date:   date,
			Open:   s.Open,
			High:   s.High,
			Low:    s.Low,
			Close:  s.Close,
			Volume: s.Volume,
		}
		msessions = append(msessions, newSession)
	}

	d.MetaData = MetaMongo(v.MetaData)
	d.Series = msessions
	d.sort()
}

func (d *TSMongo) SetM(v MonthlyTS) {
	var msessions []SessionMongo
	for date, s := range v.Series {
		newSession := SessionMongo{
			Date:   date,
			Open:   s.Open,
			High:   s.High,
			Low:    s.Low,
			Close:  s.Close,
			Volume: s.Volume,
		}
		msessions = append(msessions, newSession)
	}

	d.MetaData = MetaMongo(v.MetaData)
	d.Series = msessions
	d.sort()
}

func (d *TSMongo) Set1(v Intraday1TS) {
	var msessions []SessionMongo
	for date, s := range v.Series {
		newSession := SessionMongo{
			Date:   date,
			Open:   s.Open,
			High:   s.High,
			Low:    s.Low,
			Close:  s.Close,
			Volume: s.Volume,
		}
		msessions = append(msessions, newSession)
	}
	d.MetaData = MetaMongo(v.MetaData)
	d.Series = msessions
	d.sort()
}

func (d *TSMongo) Set5(v Intraday5TS) {
	var msessions []SessionMongo
	for date, s := range v.Series {
		newSession := SessionMongo{
			Date:   date,
			Open:   s.Open,
			High:   s.High,
			Low:    s.Low,
			Close:  s.Close,
			Volume: s.Volume,
		}
		msessions = append(msessions, newSession)
	}

	d.MetaData = MetaMongo(v.MetaData)
	d.Series = msessions
	d.sort()
}

func (d *TSMongo) Set15(v Intraday15TS) {
	var msessions []SessionMongo
	for date, s := range v.Series {
		newSession := SessionMongo{
			Date:   date,
			Open:   s.Open,
			High:   s.High,
			Low:    s.Low,
			Close:  s.Close,
			Volume: s.Volume,
		}
		msessions = append(msessions, newSession)
	}

	d.MetaData = MetaMongo(v.MetaData)
	d.Series = msessions
	d.sort()
}

func (d *TSMongo) Set30(v Intraday30TS) {
	var msessions []SessionMongo
	for date, s := range v.Series {
		newSession := SessionMongo{
			Date:   date,
			Open:   s.Open,
			High:   s.High,
			Low:    s.Low,
			Close:  s.Close,
			Volume: s.Volume,
		}
		msessions = append(msessions, newSession)
	}

	d.MetaData = MetaMongo(v.MetaData)
	d.Series = msessions
	d.sort()
}

func (d *TSMongo) Set60(v Intraday60TS) {
	var msessions []SessionMongo
	for date, s := range v.Series {

		newSession := SessionMongo{
			Date:   date,
			Open:   s.Open,
			High:   s.High,
			Low:    s.Low,
			Close:  s.Close,
			Volume: s.Volume,
		}
		msessions = append(msessions, newSession)
	}

	d.MetaData = MetaMongo(v.MetaData)
	d.Series = msessions
	d.sort()
}
