package main

import (
	"fmt"

	"github.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0.07, 0.05, 0.08, 0.15}

	for _, tax := range taxRates {
		priceJob := prices.New(tax)
		result := priceJob.Process()

		fmt.Println(result)
	}

}
