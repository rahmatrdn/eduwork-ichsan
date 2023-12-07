package main

import "fmt"

const (
	dollarToRupiah    = 15000
	maxDollarExchange = 100
)

func main() {
	saldoDollar := 1000.0
	fmt.Println("Saldo Dollar: ", saldoDollar)

	saldoRupiah := exchangeDollarToRupiah(saldoDollar)
	fmt.Println("Saldo Rupiah: ", saldoRupiah)
}

func exchangeDollarToRupiah(dollar float64) float64 {
	if dollar > maxDollarExchange {
		fmt.Println("tidak bisa menukar lebih dari ", maxDollarExchange)
		return dollar
	}

	rupiah := dollar * dollarToRupiah
	return rupiah
}
