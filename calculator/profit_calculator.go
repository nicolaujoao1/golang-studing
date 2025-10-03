package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	revenue, errorRevenue := getUserInput("Enter total revenue:")
	validate(errorRevenue)
	expenses, errExpenses := getUserInput("Enter total expenses:")
	validate(errExpenses)
	taxRate, err := getUserInput("Enter tax rate (as a decimal, e.g., 0.2 for 20%):")
	validate(err)

	ebt, profit, ratio := calculateFutureValue(revenue, expenses, taxRate)
	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.1f\n", ratio)
	storeResults(ebt, profit, ratio)

}
func storeResults(ebt, profit, ratio float64) {

	results := fmt.Sprintf("EBT: %.2f, Profit: %.2f, Ratio: %.2f\n", ebt, profit, ratio)

	os.WriteFile("results.txt", []byte(results), 0644)

}
func getUserInput(prompt string) (float64, error) {
	var input float64
	fmt.Print(prompt)
	_, err := fmt.Scan(&input)
	if err != nil {
		return 0, errors.New("invalid input, please enter a number")
	}
	if input < 0 {
		return 0, errors.New("input must be non-negative")
	}
	return input, nil
}
func validate(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
func calculateFutureValue(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}
