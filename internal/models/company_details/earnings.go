package details_models

import (
	"errors"
	"net/url"
	"stonks/internal/models"
)

type AnnualEarnings struct {
	FiscalDateEnding string `json:"fiscalDateEnding,omitempty"`
	ReportedEPS      string `json:"reportedEPS,omitempty"`
}

type QuarterlyEarnings struct {
	FiscalDateEnding   string `json:"fiscalDateEnding,omitempty"`
	ReportedDate       string `json:"reportedDate,omitempty"`
	ReportedEPS        string `json:"reportedEPS,omitempty"`
	EstimatedEPS       string `json:"estimatedEPS,omitempty"`
	Surprise           string `json:"surprise,omitempty"`
	SurprisePercentage string `json:"surprisePercentage,omitempty"`
}

type DAnnualEarnings struct {
	Symbol         string           `json:"symbol"`
	AnnualEarnings []AnnualEarnings `json:"annualEarnings"`
}

type DQuarterlyEarnings struct {
	Symbol            string              `json:"symbol"`
	QuarterlyEarnings []QuarterlyEarnings `json:"quarterlyEarnings"`
}

type Earnings struct {
	Symbol            string              `json:"symbol"`
	AnnualEarnings    []AnnualEarnings    `json:"annualEarnings"`
	QuarterlyEarnings []QuarterlyEarnings `json:"quarterlyEarnings"`
}

func (e *Earnings) Annual(t models.Timing) ([]AnnualEarnings, error) {
	if t.HasFrom() && t.HasTo() && t.HasDate() {
		return nil, errors.New("invalid data parameters")
	}

	if !t.HasFrom() && !t.HasTo() && !t.HasDate() {
		return e.AnnualEarnings, nil
	}

	var res []AnnualEarnings
	if t.HasFrom() && t.HasTo() && !t.HasDate() {
		for _, el := range e.AnnualEarnings {
			fde := el.FiscalDateEnding[:len(el.FiscalDateEnding)-6]

			if fde >= t.From && fde <= t.To {
				res = append(res, el)
			}
		}

		if len(res) == 0 {
			return nil, errors.New("no suitable data")
		}
		return res, nil
	}

	if t.HasFrom() && !t.HasTo() && !t.HasDate() {
		for _, el := range e.AnnualEarnings {
			fde := el.FiscalDateEnding[:len(el.FiscalDateEnding)-6]
			if fde >= t.From {
				res = append(res, el)
			}
		}
		if len(res) == 0 {
			return nil, errors.New("no suitable data")
		}

		return res, nil
	}

	if !t.HasFrom() && t.HasTo() && !t.HasDate() {
		for _, el := range e.AnnualEarnings {
			fde := el.FiscalDateEnding[:len(el.FiscalDateEnding)-6]
			if fde <= t.To {
				res = append(res, el)
			}
		}
		if len(res) == 0 {
			return nil, errors.New("no suitable data")
		}

		return res, nil
	}

	if t.HasDate() && !t.HasFrom() && !t.HasTo() {
		for _, el := range e.AnnualEarnings {
			fde := el.FiscalDateEnding[:len(el.FiscalDateEnding)-6]
			if fde == t.Date {
				res = append(res, el)
			}
		}
		if len(res) == 0 {
			return nil, errors.New("no suitable data")
		}

		return res, nil
	}

	return res, errors.New("bad timing")
}

func (e *Earnings) Quarterly(t models.Timing) ([]QuarterlyEarnings, error) {
	if t.HasFrom() && t.HasTo() && t.HasDate() {
		return nil, errors.New("invalid data parameters")
	}

	if !t.HasFrom() && !t.HasTo() && !t.HasDate() {
		return e.QuarterlyEarnings, nil
	}

	var res []QuarterlyEarnings
	if t.HasFrom() && t.HasTo() && !t.HasDate() {
		for _, el := range e.QuarterlyEarnings {
			fde := el.FiscalDateEnding[:len(el.FiscalDateEnding)-3]
			if fde >= t.From && fde <= t.To {
				res = append(res, el)
			}
		}
		if len(res) == 0 {
			return nil, errors.New("no suitable data")
		}

		return res, nil
	}

	if t.HasFrom() && !t.HasTo() && !t.HasDate() {
		for _, el := range e.QuarterlyEarnings {
			fde := el.FiscalDateEnding[:len(el.FiscalDateEnding)-3]
			if fde >= t.From {
				res = append(res, el)
			}
		}
		if len(res) == 0 {
			return nil, errors.New("no suitable data")
		}

		return res, nil
	}

	if !t.HasFrom() && t.HasTo() && !t.HasDate() {
		for _, el := range e.QuarterlyEarnings {
			fde := el.FiscalDateEnding[:len(el.FiscalDateEnding)-3]
			if fde <= t.To {
				res = append(res, el)
			}
		}
		if len(res) == 0 {
			return nil, errors.New("no suitable data")
		}

		return res, nil
	}

	if t.HasDate() && !t.HasFrom() && !t.HasTo() {
		for _, el := range e.QuarterlyEarnings {
			fde := el.FiscalDateEnding[:len(el.FiscalDateEnding)-3]
			if fde == t.Date {
				res = append(res, el)
			}
		}
		if len(res) == 0 {
			return nil, errors.New("no suitable data")
		}

		return res, nil
	}

	return res, errors.New("bad timing")
}

func (e *Earnings) ByTiming(values url.Values) (interface{}, error) {
	var timings models.Timing
	timings.Set(values)
	if values.Get("timing") == "annual" {
		res, err := e.Annual(timings)
		if err != nil {
			return DAnnualEarnings{Symbol: values.Get("symbol"),
				AnnualEarnings: e.AnnualEarnings}, nil
		}
		return DAnnualEarnings{Symbol: values.Get("symbol"), AnnualEarnings: res}, nil
	} else if values.Get("timing") == "quarterly" {
		res, err := e.Quarterly(timings)
		if err != nil {
			return DQuarterlyEarnings{Symbol: values.Get("symbol"),
				QuarterlyEarnings: e.QuarterlyEarnings}, nil
		}
		return DQuarterlyEarnings{Symbol: values.Get("symbol"), QuarterlyEarnings: res}, nil
	}
	return nil, errors.New("unresolved error")
}
