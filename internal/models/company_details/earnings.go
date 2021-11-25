package details_models

import (
	"strconv"
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

type Earnings struct {
	Symbol            string              `json:"symbol"`
	AnnualEarnings    []AnnualEarnings    `json:"annualEarnings" bson:"annualEarnings"`
	QuarterlyEarnings []QuarterlyEarnings `json:"quarterlyEarnings" bson:"quarterlyEarnings" `
}

type AnnualEarningsMongo struct {
	FiscalDateEnding string  `json:"fiscalDateEnding,omitempty" bson:"fiscalDateEnding,omitempty"`
	ReportedEPS      float64 `json:"reportedEPS,omitempty"      bson:"reportedEPS,omitempty"`
}

func (a *AnnualEarningsMongo) Set(v AnnualEarnings) {
	REPS, _ := strconv.ParseFloat(v.ReportedEPS, 64)

	a.FiscalDateEnding = v.FiscalDateEnding
	a.ReportedEPS = REPS
}

//TODO make a date type

type QuarterlyEarningsMongo struct {
	FiscalDateEnding   string  `json:"fiscalDateEnding,omitempty"   bson:"fiscalDateEnding"`
	ReportedDate       string  `json:"reportedDate,omitempty"       bson:"reportedDate"`
	ReportedEPS        float64 `json:"reportedEPS,omitempty"        bson:"reportedEPS"`
	EstimatedEPS       float64 `json:"estimatedEPS,omitempty"       bson:"estimatedEPS"`
	Surprise           float64 `json:"surprise,omitempty"           bson:"surprise"`
	SurprisePercentage float64 `json:"surprisePercentage,omitempty" bson:"surprisePercentage"`
}

func (q *QuarterlyEarningsMongo) Set(v QuarterlyEarnings) {
	REPS, _ := strconv.ParseFloat(v.ReportedEPS, 64)
	EEPS, _ := strconv.ParseFloat(v.EstimatedEPS, 64)
	Surprise, _ := strconv.ParseFloat(v.Surprise, 64)
	SurpriseP, _ := strconv.ParseFloat(v.SurprisePercentage, 64)

	q.FiscalDateEnding = v.FiscalDateEnding
	q.ReportedDate = v.ReportedDate
	q.ReportedEPS = REPS
	q.EstimatedEPS = EEPS
	q.Surprise = Surprise
	q.SurprisePercentage = SurpriseP
}

type EarningsMongo struct {
	Symbol            string                   `json:"symbol"              bson:"symbol"`
	AnnualEarnings    []AnnualEarningsMongo    `json:"annual,omitempty"    bson:"annual,omitempty"`
	QuarterlyEarnings []QuarterlyEarningsMongo `json:"quarterly,omitempty" bson:"quarterly,omitempty"`
}

func (e *EarningsMongo) Set(v Earnings) {
	var annual []AnnualEarningsMongo
	for _, cell := range v.AnnualEarnings {
		var NewEM AnnualEarningsMongo
		NewEM.Set(cell)

		annual = append(annual, NewEM)
	}

	var quarterly []QuarterlyEarningsMongo
	for _, cell := range v.QuarterlyEarnings {
		var NewQM QuarterlyEarningsMongo
		NewQM.Set(cell)

		quarterly = append(quarterly, NewQM)
	}

	e.Symbol = v.Symbol
	e.AnnualEarnings = annual
	e.QuarterlyEarnings = quarterly
}
