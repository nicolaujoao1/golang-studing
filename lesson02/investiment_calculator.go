package main

import (
	"fmt"
	"math"
)

func main() {

	var investmentAmount float64
	var years float64 = 10
	var expetedReturnRate float64 = 5.5
	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)
	var name string

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Printf("Hello, %s!\n", name)

	futureValue := investmentAmount * math.Pow((1+expetedReturnRate/100), years)
	fmt.Printf("Future value of the investment: %.2f\n", futureValue)
}
