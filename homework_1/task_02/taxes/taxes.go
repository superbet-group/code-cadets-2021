package taxes

import (
	"github.com/pkg/errors"
)

type TaxClass struct {
	UpperThreshold float64
	Percentage     float64
}

func validateTaxClasses(taxClasses []TaxClass) error {
	for idx, val := range taxClasses {
		if idx != len(taxClasses)-1 {
			if val.UpperThreshold >= taxClasses[idx+1].UpperThreshold || val.Percentage >= taxClasses[idx+1].Percentage {
				return errors.New("tax levels are not compatible")
			} else if val.UpperThreshold < 0 {
				return errors.New("tax levels are not compatible")
			}
		}
	}
	return nil
}

func CalculateTax(inputValue float64, taxClasses []TaxClass) (float64, error) {

	if inputValue < 0 {
		return 0, errors.New("input value is negative")
	}
	err := validateTaxClasses(taxClasses)
	if err != nil {
		return 0, err
	}

	var result float64 = 0

	for idx, class := range taxClasses {
		if class.UpperThreshold < inputValue {
			if idx != 0 {
				result += (class.UpperThreshold - taxClasses[idx-1].UpperThreshold) * class.Percentage
			} else {
				result += class.UpperThreshold * class.Percentage
			}
		} else {
			if idx != 0 {
				result += (inputValue - taxClasses[idx-1].UpperThreshold) * class.Percentage
			} else {
				result += inputValue * class.Percentage
			}
			break
		}
	}

	return result, nil
}
