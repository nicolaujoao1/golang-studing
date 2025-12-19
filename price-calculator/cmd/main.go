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
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {

		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := filemanager.New("prices/prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index])
	}

	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Printf("Processing for tax rate %.2f completed.\n", taxRates[index])
		}
	}

	// for _, doneChan := range doneChans {
	// 	<-doneChan
	// }
	// for _, errorChan := range errorChans {
	// 	select {
	// 	case err := <-errorChan:
	// 		if err != nil {
	// 			fmt.Println("Error:", err)
	// 		}
	// 	default:
	// 	}
	// }

}
