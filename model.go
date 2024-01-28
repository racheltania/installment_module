package main

import "time"

type TableInstallment struct {
	InstallmentNumber    int     `json:"installment_number"`
	Date                 string  `json:"date"`
	TotalInstallment     float64 `json:"total_installment"`
	PrincipalInstallment float64 `json:"principal_installment"`
	InterestInstallment  float64 `json:"interest_installment"`
	RemainingLoan        float64 `json:"remaining_loan"`
}

type RequestModule struct {
	Plafond             float64 `json:"plafond"`
	LoanDurationInMonth int     `json:"loan_duration_in_month"`
	AnnualInterestRate  float64 `json:"annual_interest_rate"`
	StartDate           string  `json:"start_date"`
}

type RequestData struct {
	Plafond             float64   `json:"plafond"`
	LoanDurationInMonth int       `json:"loan_duration_in_month"`
	AnnualInterestRate  float64   `json:"annual_interest_rate"`
	StartDate           time.Time `json:"start_date"`
}
