package details_models

import (
	"errors"
	"net/url"
	"stonks/internal/models"
)

type CashFlowReport struct {
	FiscalDateEnding                       string `json:"fiscalDateEnding,omitempty"`
	ReportedCurrency                       string `json:"reportedCurrency,omitempty"`
	TotalAssets                            string `json:"totalAssets,omitempty"`
	TotalCurrentAssets                     string `json:"totalCurrentAssets,omitempty"`
	CashAndCashEquivalentsAtCarryingValue  string `json:"cashAndCashEquivalentsAtCarryingValue,omitempty"`
	CashAndShortTermInvestments            string `json:"cashAndShortTermInvestments,omitempty"`
	Inventory                              string `json:"inventory,omitempty"`
	CurrentNetReceivables                  string `json:"currentNetReceivables,omitempty"`
	TotalNonCurrentAssets                  string `json:"totalNonCurrentAssets,omitempty"`
	PropertyPlantEquipment                 string `json:"propertyPlantEquipment,omitempty"`
	AccumulatedDepreciationAmortizationPPE string `json:"accumulatedDepreciationAmortizationPPE,omitempty"`
	IntangibleAssets                       string `json:"intangibleAssets,omitempty"`
	IntangibleAssetsExcludingGoodwill      string `json:"intangibleAssetsExcludingGoodwill,omitempty"`
	Goodwill                               string `json:"goodwill,omitempty"`
	Investments                            string `json:"investments,omitempty"`
	LongTermInvestments                    string `json:"longTermInvestments,omitempty"`
	ShortTermInvestments                   string `json:"shortTermInvestments,omitempty"`
	OtherCurrentAssets                     string `json:"otherCurrentAssets,omitempty"`
	OtherNonCurrrentAssets                 string `json:"otherNonCurrrentAssets,omitempty"`
	TotalLiabilities                       string `json:"totalLiabilities,omitempty"`
	TotalCurrentLiabilities                string `json:"totalCurrentLiabilities,omitempty"`
	CurrentAccountsPayable                 string `json:"currentAccountsPayable,omitempty"`
	DeferredRevenue                        string `json:"deferredRevenue,omitempty"`
	CurrentDebt                            string `json:"currentDebt,omitempty"`
	ShortTermDebt                          string `json:"shortTermDebt,omitempty"`
	TotalNonCurrentLiabilities             string `json:"totalNonCurrentLiabilities,omitempty"`
	CapitalLeaseObligations                string `json:"capitalLeaseObligations,omitempty"`
	LongTermDebt                           string `json:"longTermDebt,omitempty"`
	CurrentLongTermDebt                    string `json:"currentLongTermDebt,omitempty"`
	LongTermDebtNoncurrent                 string `json:"longTermDebtNoncurrent,omitempty"`
	ShortLongTermDebtTotal                 string `json:"shortLongTermDebtTotal,omitempty"`
	OtherCurrentLiabilities                string `json:"otherCurrentLiabilities,omitempty"`
	OtherNonCurrentLiabilities             string `json:"otherNonCurrentLiabilities,omitempty"`
	TotalShareholderEquity                 string `json:"totalShareholderEquity,omitempty"`
	TreasuryStock                          string `json:"treasuryStock,omitempty"`
	RetainedEarnings                       string `json:"retainedEarnings,omitempty"`
	CommonStock                            string `json:"commonStock,omitempty"`
	CommonStockSharesOutstanding           string `json:"commonStockSharesOutstanding,omitempty"`
}

type DAnnualCashFlowReports struct {
	Symbol         string           `json:"symbol"`
	AnnualEarnings []CashFlowReport `json:"annualReports"`
}

type DQuarterlyCashFlowReports struct {
	Symbol            string           `json:"symbol"`
	QuarterlyEarnings []CashFlowReport `json:"quarterlyReports"`
}

type CashFlow struct {
	Symbol           string           `json:"symbol"`
	AnnualReports    []CashFlowReport `json:"annualReports"`
	QuarterlyReports []CashFlowReport `json:"quarterlyReports"`
}

func (c *CashFlow) Annual(t models.Timing) ([]CashFlowReport, error) {
	if t.HasFrom() && t.HasTo() && t.HasDate() {
		return nil, errors.New("invalid data parameters")
	}

	if !t.HasFrom() && !t.HasTo() && !t.HasDate() {
		return c.QuarterlyReports, nil
	}

	var res []CashFlowReport
	if t.HasFrom() && t.HasTo() && !t.HasDate() {
		for _, el := range c.AnnualReports {
			fde := el.FiscalDateEnding[:len(el.FiscalDateEnding)-5]

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
		for _, el := range c.AnnualReports {
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
		for _, el := range c.AnnualReports {
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
		for _, el := range c.AnnualReports {
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

func (c *CashFlow) Quarterly(t models.Timing) ([]CashFlowReport, error) {
	if t.HasFrom() && t.HasTo() && t.HasDate() {
		return nil, errors.New("invalid data parameters")
	}

	if !t.HasFrom() && !t.HasTo() && !t.HasDate() {
		return c.QuarterlyReports, nil
	}

	var res []CashFlowReport
	if t.HasFrom() && t.HasTo() && !t.HasDate() {
		for _, el := range c.QuarterlyReports {
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
		for _, el := range c.QuarterlyReports {
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
		for _, el := range c.QuarterlyReports {
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
		for _, el := range c.QuarterlyReports {
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

func (c *CashFlow) ByTiming(values url.Values) (interface{}, error) {
	var timings models.Timing
	timings.Set(values)
	if values.Get("timing") == "annual" {
		res, err := c.Annual(timings)
		if err != nil {
			return DAnnualCashFlowReports{Symbol: values.Get("symbol"),
				AnnualEarnings: c.AnnualReports}, nil
		}
		return DAnnualCashFlowReports{Symbol: values.Get("symbol"), AnnualEarnings: res}, nil
	} else if values.Get("timing") == "quarterly" {
		res, err := c.Quarterly(timings)
		if err != nil {
			return DQuarterlyCashFlowReports{Symbol: values.Get("symbol"),
				QuarterlyEarnings: c.QuarterlyReports}, nil
		}
		return DQuarterlyCashFlowReports{Symbol: values.Get("symbol"), QuarterlyEarnings: res}, nil
	}
	return nil, errors.New("unresolved error")
}
