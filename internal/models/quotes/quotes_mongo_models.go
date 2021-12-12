package quotes_models

import (
	"sort"
	"strconv"
)

type SessionMongo struct {
	Date   string  `json:"date,omitempty"`
	Open   float64 `json:"open,omitempty"`
	High   float64 `json:"high,omitempty"`
	Low    float64 `json:"low,omitempty"`
	Close  float64 `json:"close,omitempty"`
	Volume float64 `json:"volume,omitempty"`
}

type MetaMongo struct {
	Information   string `json:"information,omitempty" bson:"information,omitempty"`
	Symbol        string `json:"symbol,omitempty" bson:"symbol,omitempty"`
	LastRefreshed string `json:"lastRefreshed,omitempty" bson:"lastRefreshed,omitempty"`
	TimeZone      string `json:"timeZone,omitempty" bson:"timeZone,omitempty"`
}

type TSMongo struct {
	MetaData MetaMongo      `json:"_meta,omitempty" bson:"_meta,omitempty"`
	Series   []SessionMongo `json:"series,omitempty"`
}

func (d *TSMongo) sort() {
	sort.Slice(d.Series, func(i, j int) bool { return d.Series[i].Date > d.Series[j].Date })
}

func newSession(v Session, d string) SessionMongo {
	open, _ := strconv.ParseFloat(v.Open, 32)
	high, _ := strconv.ParseFloat(v.High, 32)
	low, _ := strconv.ParseFloat(v.Low, 32)
	closePrice, _ := strconv.ParseFloat(v.Close, 32)
	volume, _ := strconv.ParseFloat(v.Volume, 32)

	ns := SessionMongo{
		Date:   d,
		Open:   open,
		High:   high,
		Low:    low,
		Close:  closePrice,
		Volume: volume,
	}

	return ns
}

func (d *TSMongo) SetD(v DailyTS) {
	var msessions []SessionMongo
	for date, s := range v.Series {
		ns := newSession(s, date)
		msessions = append(msessions, ns)
	}

	d.MetaData = MetaMongo(v.MetaData)
	d.Series = msessions
	d.sort()
}

func (d *TSMongo) SetW(v WeeklyTS) {
	var msessions []SessionMongo
	for date, s := range v.Series {
		ns := newSession(s, date)
		msessions = append(msessions, ns)
	}

	d.MetaData = MetaMongo(v.MetaData)
	d.Series = msessions
	d.sort()
}

func (d *TSMongo) SetM(v MonthlyTS) {
	var msessions []SessionMongo
	for date, s := range v.Series {
		ns := newSession(s, date)
		msessions = append(msessions, ns)
	}

	d.MetaData = MetaMongo(v.MetaData)
	d.Series = msessions
	d.sort()
}

func (d *TSMongo) Set1(v Intraday1TS) {
	var msessions []SessionMongo
	for date, s := range v.Series {
		ns := newSession(s, date)
		msessions = append(msessions, ns)
	}
	d.MetaData = MetaMongo(v.MetaData)
	d.Series = msessions
	d.sort()
}

func (d *TSMongo) Set5(v Intraday5TS) {
	var msessions []SessionMongo
	for date, s := range v.Series {
		ns := newSession(s, date)
		msessions = append(msessions, ns)
	}

	d.MetaData = MetaMongo(v.MetaData)
	d.Series = msessions
	d.sort()
}

func (d *TSMongo) Set15(v Intraday15TS) {
	var msessions []SessionMongo
	for date, s := range v.Series {
		ns := newSession(s, date)
		msessions = append(msessions, ns)
	}

	d.MetaData = MetaMongo(v.MetaData)
	d.Series = msessions
	d.sort()
}

func (d *TSMongo) Set30(v Intraday30TS) {
	var msessions []SessionMongo
	for date, s := range v.Series {
		ns := newSession(s, date)
		msessions = append(msessions, ns)
	}

	d.MetaData = MetaMongo(v.MetaData)
	d.Series = msessions
	d.sort()
}

func (d *TSMongo) Set60(v Intraday60TS) {
	var msessions []SessionMongo
	for date, s := range v.Series {
		ns := newSession(s, date)
		msessions = append(msessions, ns)
	}

	d.MetaData = MetaMongo(v.MetaData)
	d.Series = msessions
	d.sort()
}
