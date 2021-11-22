package models

import (
	"net/url"
	"strconv"
)

type PriceTag struct {
	Min float64
	Max float64
}

func (t *PriceTag) Set(v url.Values) {
	min, _ := strconv.ParseFloat(v.Get("min"), 32)
	max, _ := strconv.ParseFloat(v.Get("max"), 32)
	t.Min = min
	t.Max = max
}

func (t *PriceTag) Has(field string) bool {
	switch field {
	case "min":
		return t.Min != 0
	case "max":
		return t.Max != 0
	default:
		return false
	}
}

func (t *PriceTag) Get(field string) float64 {
	switch field {
	case "min":
		return t.Min
	case "max":
		return t.Max
	default:
		return 0
	}
}
