package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"math"

	"code-cadets-2021/homework_1/task_02/taxes"
)

func getIncome()(float64, error){
	var income float64
	fmt.Sprintf("Enter income value: ")

	_, err := fmt.Scanf("%f", &income)
	if err != nil {
		return 0, errors.New("invalid value")
	}
	return income, nil
}

func main() {

	income, err := getIncome()
	if err != nil {
		log.Fatal(err)
	}

	percentages := []float64{0.0, 0.10, 0.20, 0.30}
	amounts := []float64{1000, 5000, 10000, math.Inf(1)}

	tax, err := taxes.CalculateTax(income, percentages, amounts)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "trouble calculating tax"),
		)
	}

	fmt.Printf("Total taxes for the given income %v equal %v.", income, tax)

}