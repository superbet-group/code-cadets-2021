package taxes

import "github.com/pkg/errors"

type taxBrackets struct {
	brackets []float32
	taxes    []float32
}

func checkTaxBrackets(brackets []float32, taxes []float32) error {
	if len(brackets) == 0 || brackets[0] != 0 {
		return errors.New("First bracket is not zero.")
	}

	if len(brackets) != len(taxes) {
		return errors.New("Brackets and taxes don't have the same length.")
	}

	for idx, bracket := range brackets {
		if idx > 0 && bracket <= brackets[idx - 1] {
			return errors.New("Bracket is not greater than the previous.")
		}

		if taxes[idx] < 0 {
			return errors.New("Tax is negative.")
		}
	}

	return nil
}

func NewTaxBrackets(brackets []float32, taxes []float32) (*taxBrackets, error) {
	if err := checkTaxBrackets(brackets, taxes); err != nil {
		return nil, err
	}

	return &taxBrackets{brackets: brackets, taxes: taxes}, nil
}

func CalculateProgressiveTax(value float32, taxBrackets *taxBrackets) (float32, error) {
	if value < 0 {
		return 0, errors.New("Value is negative.")
	}

	brackets := taxBrackets.brackets
	taxes := taxBrackets.taxes

	var finalTax float32

	for idx, tax := range taxes {
		if idx + 1 == len(taxes) || value <= brackets[idx + 1] {
			finalTax += (value - brackets[idx]) * tax
			break
		}

		finalTax += (brackets[idx + 1] - brackets[idx]) * tax
	}

	return finalTax, nil
}