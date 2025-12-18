package main

import (
	// "github.com/price-calculator/filemanager"
	// "github.com/price-calculator/cmdmanager"
	"fmt"
	// "github.com/price-calculator/cmdmanager"
	"github.com/price-calculator/filemanager"
	"github.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices/prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("Could not process Job")
			fmt.Println("Error:", err)
		}
	}

}
