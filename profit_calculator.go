package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	// fmt.Print("Revenue($): ")
	// fmt.Scan(&revenue)
	// fmt.Print("Expenses($): ")
	// fmt.Scan(&expenses)
	// fmt.Print("Tax Rate(%): ")
	// fmt.Scan(&taxRate)
	revenue, err := getUserValue("Revenue($): ")
	if err != nil {
		handleInvalidValue(err)
	}
	expenses, err := getUserValue("Expenses($): ")
	if err != nil {
		handleInvalidValue(err)
	}
	taxRate, err := getUserValue("Tax Rate(%): ")
	if err != nil {
		handleInvalidValue(err)
	}

	// earningsBeforeTax := revenue - expenses
	// profit := earningsBeforeTax * (1 - taxRate/100) // Earnings After Tax
	// ratio := earningsBeforeTax / profit
	earningsBeforeTax, profit, ratio := doMaths(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", earningsBeforeTax)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f", ratio)
	storeResults(earningsBeforeTax, profit, ratio)

}

func getUserValue(infoText string) (float64, error) {
	var userValue float64
	fmt.Print(infoText)
	fmt.Scan(&userValue)
	if userValue <= 0 {
		return 0, errors.New("the value should not be zero or negative")
	}
	return userValue, nil
}

func handleInvalidValue(err error) {
	fmt.Println("-------")
	fmt.Println("ERROR")
	fmt.Println(err)
	panic("Can't continue, sorry.")
}

func doMaths(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100) // Earnings After Tax
	ratio := earningsBeforeTax / profit
	return earningsBeforeTax, profit, ratio
}

func storeResults(earningsBeforeTax, profit, ratio float64) {
	valueText := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", earningsBeforeTax, profit, ratio)
	os.WriteFile("values.txt", []byte(valueText), 0644)
}
