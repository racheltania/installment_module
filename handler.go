package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func HandlerRequestInstallment(ctx echo.Context) error {
	requestModule := new(RequestModule)
	if err := ctx.Bind(&requestModule); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}
	parsedDate, errParsed := time.Parse("2006-01-02", requestModule.StartDate)
	if errParsed != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid request",
			"details": errParsed.Error(),
		})
	}

	requestData := RequestData{
		Plafond:             requestModule.Plafond,
		LoanDurationInMonth: requestModule.LoanDurationInMonth,
		AnnualInterestRate:  requestModule.AnnualInterestRate,
		StartDate:           parsedDate,
	}

	result := installmentModule(requestData)
	return ctx.JSON(http.StatusOK, result)
}
