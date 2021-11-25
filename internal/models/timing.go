package models

import "net/url"

type Timing struct {
	Symbol   string
	From     string
	To       string
	Date     string
	Interval string
}

func (t *Timing) Set(v url.Values) {
	t.Symbol = v.Get("symbol")
	t.From = v.Get("from")
	t.To = v.Get("to")
	t.Date = v.Get("date")
	t.Interval = v.Get("interval")
}

func (t *Timing) Has(field string) bool {
	switch field {
	case "symbol":
		return t.Symbol != ""
	case "from":
		return t.From != ""
	case "to":
		return t.To != ""
	case "date":
		return t.Date != ""
	case "interval":
		return t.Interval != ""
	default:
		return false
	}
}

func (t *Timing) HasInterval() bool {
	return t.Has("from") || t.Has("to") || t.Has("date")
}

func (t *Timing) Get(field string) string {
	switch field {
	case "symbol":
		return t.Symbol
	case "from":
		return t.From
	case "to":
		return t.To
	case "date":
		return t.Date
	case "interval":
		return t.Interval
	default:
		return ""
	}
}
