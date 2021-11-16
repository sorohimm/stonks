package models

import "net/url"

type Timing struct {
	Symbol string
	From   string
	To     string
	Date   string
}

func (t *Timing) Set(v url.Values) {
	t.Symbol = v.Get("symbol")
	t.From = v.Get("from")
	t.To = v.Get("to")
	t.Date = v.Get("date")
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
	default:
		return false
	}
}


func (t *Timing) HasFrom() bool {
	return t.From != ""
}

func (t *Timing) HasTo() bool {
	return t.To != ""
}

func (t *Timing) HasDate() bool {
	return t.Date != ""
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
	default:
		return ""
	}
}
