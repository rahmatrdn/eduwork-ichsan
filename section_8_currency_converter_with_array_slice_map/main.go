package main

import (
	"fmt"
	"time"
)

var exchangeMap = make(map[string]float64)

type Request struct {
	amount       float64
	currencyFrom string
	currencyTo   string
}

func CurrencyConversion(req Request) (float64, string) {
	if req.amount <= 0 {
		return 0, "invalid amount"
	}

	exchangeValue := exchangeMap[req.currencyTo]
	if exchangeValue == 0 || req.currencyFrom != "USD" {
		return 0, "your currency is invalid, check available currency"
	}

	return req.amount * exchangeValue, ""
}

func main() {
	// start date exchange value
	startDate := time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC)
	today := time.Now()
	duration := today.Sub(startDate)
	durationDay := int(duration.Hours() / 24)

	currencyNames := [...]string{"EUR", "GBP", "JPY"}
	exchangeValues := []float64{
		0.90,
		0.80,
		145.50,
	}

	for i := 0; i < len(currencyNames); i++ {
		if duration > 0 {
			// multiply the exchangeValue per day
			exchangeMap[currencyNames[i]] = exchangeValues[i] * (float64(durationDay) + 1)
		} else {
			exchangeMap[currencyNames[i]] = exchangeValues[i]
		}
	}

	req := HandlerRequest()
	res, err := CurrencyConversion(req)
	if err != "" {
		fmt.Println(err)
	} else {
		fmt.Println("Result: ", res)
	}
}

func HandlerRequest() Request {
	var req Request

	fmt.Print("Input Money: ")
	fmt.Scanln(&req.amount)

	fmt.Print("Input Currency From (Only USD): ")
	fmt.Scanln(&req.currencyFrom)

	fmt.Print("Input Currency To: ")
	fmt.Scanln(&req.currencyTo)

	return req
}
