package main

import (
	"fmt"

	"example.com/price-calculator/filemanger"
	"example.com/price-calculator/prices"
)

func main() {

	taxRate := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRate {
		fm := filemanger.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		priceJob.Process()
	}
}
