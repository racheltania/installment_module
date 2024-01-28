package main

import (
	"math"
)

func totalInstallment(plafond float64, loanDurationInMonth int, annualInterestRate float64) float64 {
	monthlyInterestRate := (annualInterestRate / 12) / 100
	formula := math.Pow(1+monthlyInterestRate, float64(loanDurationInMonth))
	result := monthlyInterestRate * plafond * (formula / (formula - 1))
	return result
}

func interestInstallment(annualInterestRate float64, remainingLoan float64) float64 {
	annualInterestRate = annualInterestRate / 100
	result := (annualInterestRate / 360) * 30 * remainingLoan
	return result
}

func roundToTwoDecimalPlaces(number float64) float64 {
	return math.Round(number*100) / 100
}

func installmentModule(request RequestData) []TableInstallment {
	var results []TableInstallment
	remainingLoan := request.Plafond
	currentDate := request.StartDate.AddDate(0, -1, 0)
	for i := 1; i <= request.LoanDurationInMonth; i++ {
		resultTotalInstallment := totalInstallment(request.Plafond, request.LoanDurationInMonth, request.AnnualInterestRate)
		currentInterestInstallment := interestInstallment(request.AnnualInterestRate, remainingLoan)
		currentPrincipalInstallment := resultTotalInstallment - currentInterestInstallment
		remainingLoan -= currentPrincipalInstallment
		currentDate = currentDate.AddDate(0, 1, 0)
		result := TableInstallment{
			InstallmentNumber:    i,
			Date:                 currentDate.Format("2006-01-02"),
			TotalInstallment:     roundToTwoDecimalPlaces(resultTotalInstallment),
			PrincipalInstallment: roundToTwoDecimalPlaces(currentPrincipalInstallment),
			InterestInstallment:  roundToTwoDecimalPlaces(currentInterestInstallment),
			RemainingLoan:        roundToTwoDecimalPlaces(remainingLoan),
		}
		results = append(results, result)
	}
	return results
}
