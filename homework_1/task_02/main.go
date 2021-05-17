package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"math"

	"code-cadets-2021/homework_1/task_02/taxes"
)

func getIncome() (float64, error) {
	var income float64
	fmt.Sprintf("Enter income value: ")

	_, err := fmt.Scanf("%f", &income)
	if err != nil {
		return 0, errors.New("invalid value")
	}
	return income, nil
}

func buildTaxClasses(amounts, percentages []float64) []taxes.TaxClass {
	var classes []taxes.TaxClass

	for i := 0; i < len(percentages); i++ {
		tempClass := taxes.TaxClass{
			UpperThreshold: amounts[i],
			Percentage:     percentages[i],
		}
		classes = append(classes, tempClass)
	}
	return classes
}

func main() {

	income, err := getIncome()
	if err != nil {
		log.Fatal(err)
	}

	percentages := []float64{0.0, 0.10, 0.20, 0.30}
	amounts := []float64{1000, 5000, 10000, math.Inf(1)}
	classes := buildTaxClasses(amounts, percentages)
	tax, err := taxes.CalculateTax(income, classes)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "trouble calculating tax"),
		)
	}

	fmt.Printf("Total taxes for the given income %v equal %v.", income, tax)

}
