package models

import (
	"net/url"
	"strconv"
)

type ChooseTag struct {
	Min   float64
	Max   float64
	Point string
}

func (t *ChooseTag) Set(v url.Values) {
	min, _ := strconv.ParseFloat(v.Get("min"), 32)
	max, _ := strconv.ParseFloat(v.Get("max"), 32)
	t.Min = min
	t.Max = max
	t.Point = v.Get("point")
}

func (t *ChooseTag) Has(field string) bool {
	switch field {
	case "min":
		return t.Min != 0
	case "max":
		return t.Max != 0
	case "point":
		return t.Point != ""
	default:
		return false
	}
}

func (t *ChooseTag) Get(field string) interface{} {
	switch field {
	case "min":
		return t.Min
	case "max":
		return t.Max
	case "point":
		return t.Point
	default:
		return 0
	}
}
