package main

import (
	"fmt"

	"code-cadets-2021/homework_1/02_ProgressiveTax/taxes"
)

func main() {
	bracketValues := []float32{0, 1000, 5000, 10000}
	taxValues := []float32{0, 0.1, 0.2, 0.3}
	var value float32 = 7000

	taxBrackets, err := taxes.NewTaxBrackets(bracketValues, taxValues)

	if err == nil {
		tax, _ := taxes.CalculateProgressiveTax(value, taxBrackets)
		fmt.Println(tax)
	}
}
