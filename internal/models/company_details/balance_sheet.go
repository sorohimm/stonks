package details_models

import (
	"errors"
	"net/url"
	"stonks/internal/models"
)

type BalanceSheetReports struct {
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

type DAnnualBalanceSheetReports struct {
	Symbol         string                `json:"symbol"`
	AnnualEarnings []BalanceSheetReports `json:"annualReports"`
}

type DQuarterlyBalanceSheetReports struct {
	Symbol            string                `json:"symbol"`
	QuarterlyEarnings []BalanceSheetReports `json:"quarterlyReports"`
}

type BalanceSheet struct {
	Symbol           string                `json:"symbol"`
	AnnualReports    []BalanceSheetReports `json:"annualReports"`
	QuarterlyReports []BalanceSheetReports `json:"quarterlyReports"`
}

func (b *BalanceSheet) Annual(t models.Timing) ([]BalanceSheetReports, error) {
	if t.HasFrom() && t.HasTo() && t.HasDate() {
		return nil, errors.New("invalid data parameters")
	}

	if !t.HasFrom() && !t.HasTo() && !t.HasDate() {
		return b.AnnualReports, nil
	}

	var res []BalanceSheetReports
	if t.HasFrom() && t.HasTo() && !t.HasDate() {
		for _, el := range b.AnnualReports {
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
		for _, el := range b.AnnualReports {
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
		for _, el := range b.AnnualReports {
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
		for _, el := range b.AnnualReports {
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

func (b *BalanceSheet) Quarterly(t models.Timing) ([]BalanceSheetReports, error) {
	if t.HasFrom() && t.HasTo() && t.HasDate() {
		return nil, errors.New("invalid data parameters")
	}

	if !t.HasFrom() && !t.HasTo() && !t.HasDate() {
		return b.QuarterlyReports, nil
	}

	var res []BalanceSheetReports
	if t.HasFrom() && t.HasTo() && !t.HasDate() {
		for _, el := range b.QuarterlyReports {
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
		for _, el := range b.QuarterlyReports {
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
		for _, el := range b.QuarterlyReports {
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
		for _, el := range b.QuarterlyReports {
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

func (b *BalanceSheet) ByTiming(values url.Values) (interface{}, error) {
	var timings models.Timing
	timings.Set(values)
	if values.Get("timing") == "annual" {
		res, err := b.Annual(timings)
		if err != nil {
			return DAnnualBalanceSheetReports{Symbol: values.Get("symbol"),
				AnnualEarnings: b.AnnualReports}, nil
		}
		return DAnnualBalanceSheetReports{Symbol: values.Get("symbol"), AnnualEarnings: res}, nil
	} else if values.Get("timing") == "quarterly" {
		res, err := b.Quarterly(timings)
		if err != nil {
			return DQuarterlyBalanceSheetReports{Symbol: values.Get("symbol"),
				QuarterlyEarnings: b.QuarterlyReports}, nil
		}
		return DQuarterlyBalanceSheetReports{Symbol: values.Get("symbol"), QuarterlyEarnings: res}, nil
	}
	return nil, errors.New("unresolved error")
}
