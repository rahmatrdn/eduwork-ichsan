package main

import "fmt"

type Request struct {
	amount       float64
	currencyFrom string
	currencyTo   string
	exchangeRate float64
}

func (req Request) ConvertCurrency() float64 {
	conversion := req.amount * 15000
	res := conversion + (conversion * req.exchangeRate / 100)
	return res
}

func Service(r Request) (float64, string) {
	if r.currencyFrom != "USD" || r.currencyTo != "RP" {
		return 0.0, "your currency cannot be converted, check available currency"
	}

	res := r.ConvertCurrency()
	return res, ""
}

func main() {
	req := HandlerRequest()
	res, err := Service(req)
	if err != "" {
		fmt.Print(err)
	} else {
		fmt.Println("Result: ", res)
	}
}

func HandlerRequest() Request {
	var req Request

	fmt.Print("Input Money: ")
	fmt.Scanln(&req.amount)

	fmt.Print("Input Currency From: ")
	fmt.Scanln(&req.currencyFrom)

	fmt.Print("Input Currency To: ")
	fmt.Scanln(&req.currencyTo)

	fmt.Print("Input Exchange Rate: ")
	fmt.Scanln(&req.exchangeRate)

	return req
}
