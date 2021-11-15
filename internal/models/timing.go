package models

import "net/url"

type Timing struct {
	Symbol string
	From   string
	To     string
	Date   string
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

func (t *Timing) Set(v url.Values) {
	t.Symbol = v.Get("symbol")
	t.From = v.Get("from")
	t.To = v.Get("to")
	t.Date = v.Get("date")
}

func (t *Timing) Get(s string) string {
	switch s {
	case "symbol":
		return t.To
	case "from":
		return t.To
	case "to":
		return t.To
	case "date":
		return t.To
	}
	return ""
}
