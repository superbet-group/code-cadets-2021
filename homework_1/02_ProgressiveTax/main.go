package main

import (
	"fmt"
	"math"
	"os"

	"code-cadets-2021/homework_1/02_ProgressiveTax/taxes"
)

func main() {
	thresholds := []float32{10000, 20000, math.MaxFloat32}
	taxRates := []float32{0.1, 0.2, 0.3}
	var value float32 = 20000

	taxBrackets, err := taxes.CreateBrackets(thresholds, taxRates)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tax, err := taxes.CalculateProgressiveTax(value, taxBrackets)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(tax)
}
