package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue($): ")
	fmt.Scan(&revenue)
	fmt.Print("Expenses($): ")
	fmt.Scan(&expenses)
	fmt.Print("Tax Rate(%): ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100) // Earnings After Tax
	ratio := earningsBeforeTax / profit

	fmt.Println(earningsBeforeTax)
	fmt.Println(profit)
	fmt.Println(ratio)
}
