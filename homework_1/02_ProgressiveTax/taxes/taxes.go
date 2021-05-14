package taxes

import (
	"sort"

	"github.com/pkg/errors"
)

type bracket struct {
	threshold float32
	taxRate   float32
}

func newBracket(threshold, taxRate float32) (*bracket, error) {
	if threshold < 0 {
		return nil, errors.New("bracket threshold is negative")
	}

	if taxRate < 0 {
		return nil, errors.New("bracket tax rate is negative")
	}

	return &bracket{threshold: threshold, taxRate: taxRate}, nil
}

func CreateBrackets(thresholds, taxRates []float32) ([]*bracket, error) {
	if len(thresholds) != len(taxRates) {
		return nil, errors.New("thresholds and taxRates do not have the same length")
	}

	var brackets []*bracket

	for idx, threshold := range thresholds {
		bracket, err := newBracket(threshold, taxRates[idx])

		if err != nil {
			return nil, err
		}

		brackets = append(brackets, bracket)
	}

	return brackets, nil
}

func CalculateProgressiveTax(value float32, brackets []*bracket) (float32, error) {
	if value < 0 {
		return 0, errors.New("value is negative")
	}

	if len(brackets) == 0 {
		return 0, nil
	}

	sort.Slice(brackets, func(i, j int) bool {
		return brackets[i].threshold < brackets[j].threshold
	})

	if value <= brackets[0].threshold {
		return value * brackets[0].taxRate, nil
	}

	tax := brackets[0].threshold * brackets[0].taxRate

	for idx, bracket := range brackets[1:] {
		if value <= bracket.threshold {
			tax += (value - brackets[idx].threshold) * bracket.taxRate
			break
		}

		tax += (bracket.threshold - brackets[idx].threshold) * bracket.taxRate
	}

	return tax, nil
}
