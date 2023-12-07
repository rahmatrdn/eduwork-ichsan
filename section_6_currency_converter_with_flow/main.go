package main

import (
	"fmt"
)

type Conversion struct {
	money        float64
	currencyFrom string
	currencyTo   string
}

func CheckCurrencyFrom(currencyFrom string) bool {
	switch currencyFrom {
	case "USD":
		return true
	case "EUR":
		return true
	case "GBP":
		return true
	case "JPY":
		return true
	}

	return false
}

func CheckCurrencyTo(currencyTo string) bool {
	switch currencyTo {
	case "USD":
		return true
	case "EUR":
		return true
	case "GBP":
		return true
	case "JPY":
		return true
	}

	return false
}

func (c Conversion) CurrencyConversion() float64 {
	total := c.money
	currencyFrom := c.currencyFrom
	currencyTo := c.currencyTo

	if currencyFrom == "USD" {
		if currencyTo == "EUR" {
			total *= 0.91
		} else if currencyTo == "GBP" {
			total *= 0.79
		} else if currencyTo == "JPY" {
			total *= 149.47
		}
	} else if currencyFrom == "EUR" {
		if currencyTo == "USD" {
			total *= 1.09
		} else if currencyTo == "GBP" {
			total *= 0.87
		} else if currencyTo == "JPY" {
			total *= 163.65
		}
	} else if currencyFrom == "GBP" {
		if currencyTo == "USD" {
			total *= 1.26
		} else if currencyTo == "EUR" {
			total *= 1.15
		} else if currencyTo == "JPY" {
			total *= 188.34
		}
	} else {
		if currencyTo == "USD" {
			total *= 0.0067
		} else if currencyTo == "EUR" {
			total *= 0.0061
		} else if currencyTo == "GBP" {
			total *= 0.0053
		}
	}

	return total
}

func Service(c Conversion) (float64, string) {
	if c.currencyFrom == c.currencyTo {
		return 0.0, "cannot convert the same value"
	}

	validCurrencyFrom := CheckCurrencyFrom(c.currencyFrom)
	validCurrencyTo := CheckCurrencyTo(c.currencyTo)
	if !validCurrencyFrom || !validCurrencyTo {
		return 0.0, "your currency cannot be converted, check available currency"
	}

	res := c.CurrencyConversion()
	if res > 1000 {
		return 0.0, "cannot convert large amounts"
	}
	return res, ""
}

func main() {
	for {
		var repeat string
		req := HandlerRequest()

		res, err := Service(req)
		if err != "" {
			fmt.Print(err, ", continue or stop? ")
			fmt.Scanln(&repeat)

			if repeat == "continue" {
				continue
			} else {
				break
			}
		} else {
			fmt.Println(res)
			break
		}
	}
}

func HandlerRequest() Conversion {
	var req Conversion

	fmt.Print("Input Money: ")
	fmt.Scanln(&req.money)

	fmt.Print("Input Currency From: ")
	fmt.Scanln(&req.currencyFrom)

	fmt.Print("Input Currency To: ")
	fmt.Scanln(&req.currencyTo)

	return req
}
