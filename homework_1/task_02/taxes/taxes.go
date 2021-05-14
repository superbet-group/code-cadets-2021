package taxes

import (
	"github.com/pkg/errors"
)

type TaxClass struct {
	UpperAmount  float64
	Percentage float64
}


func checkLevels(taxClasses []TaxClass) bool {
	for idx, val := range taxClasses {
		if idx != len(taxClasses)-1 {
			if val.UpperAmount >= taxClasses[idx+1].UpperAmount || val.Percentage >= taxClasses[idx+1].Percentage {
				return false
			}
		}
	}
	return true
}

func buildTaxClasses(amounts, percentages []float64) []TaxClass{
	var classes []TaxClass

	for i :=0; i < len(percentages); i++{
		tempClass := TaxClass{
			UpperAmount: amounts[i],
			Percentage: percentages[i],
		}
		classes = append(classes, tempClass)
	}
	return classes
}

func CalculateTax(inputValue float64, percentages, amounts []float64) (float64, error){

	if inputValue < 0 {
		return 0, errors.New("input value is negative")
	}
	taxClasses := buildTaxClasses(amounts, percentages)
	var okLevels bool = checkLevels(taxClasses)

	if !okLevels {
		return 0, errors.New("tax levels are not compatible")
	}

	var result float64 = 0

	for idx, class := range taxClasses {
		if class.UpperAmount < inputValue {
			if idx != 0 {
				result += (class.UpperAmount-taxClasses[idx-1].UpperAmount)*class.Percentage
			} else {
				result += class.UpperAmount * class.Percentage
			}
		} else {
			if idx != 0 {
				result += (inputValue - taxClasses[idx-1].UpperAmount) * class.Percentage
			} else {
				result += inputValue * class.Percentage
			}
			break
		}
	}

	return result, nil
}