package main

import (
	"fmt"
)

func main() {

	// fmt.Print("Revenue($): ")
	// fmt.Scan(&revenue)
	// fmt.Print("Expenses($): ")
	// fmt.Scan(&expenses)
	// fmt.Print("Tax Rate(%): ")
	// fmt.Scan(&taxRate)
	revenue := getUserValue("Revenue($): ")
	expenses := getUserValue("Expenses($): ")
	taxRate := getUserValue("Tax Rate(%): ")

	// earningsBeforeTax := revenue - expenses
	// profit := earningsBeforeTax * (1 - taxRate/100) // Earnings After Tax
	// ratio := earningsBeforeTax / profit
	earningsBeforeTax, profit, ratio := doMaths(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", earningsBeforeTax)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f", ratio)
}

func getUserValue(infoText string) float64 {
	var userValue float64
	fmt.Print(infoText)
	fmt.Scan(&userValue)
	return userValue
}

func doMaths(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100) // Earnings After Tax
	ratio := earningsBeforeTax / profit
	return earningsBeforeTax, profit, ratio
}
