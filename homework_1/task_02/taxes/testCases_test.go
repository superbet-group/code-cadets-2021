package taxes_test

import "math"

type testCase struct {
	inputAmount float64
	percentages []float64
	amounts []float64
	outputTax float64
	error bool
}

func getTestCases() []testCase {
	return []testCase {
		{
			inputAmount: 7000,
			percentages: []float64{0.0, 0.10, 0.20, 0.30},
			amounts: []float64{1000, 5000, 10000, math.Inf(1)},
			outputTax: 800,
			error: false,
		},
		{
			inputAmount: 456456,
			percentages: []float64{0.0, 0.10, 0.20, 0.30},
			amounts: []float64{1000, 5000, 10000, math.Inf(1)},
			outputTax: 135336.8,
			error: false,
		},
		{
			inputAmount: -2500,
			percentages: []float64{0.0, 0.10, 0.20, 0.30},
			amounts: []float64{1000, 5000, 10000, math.Inf(1)},
			error: true,
		},
		{
			inputAmount: 123123,
			percentages: []float64{0.0, 0.10, 0.20, 0.30},
			amounts: []float64{1000, 5000, 10000, math.Inf(1)},
			outputTax: 35336.9,
			error: false,
		},
		{
			inputAmount: 1000,
			percentages: []float64{0.0, 0.10, 0.20, 0.30},
			amounts: []float64{1000, 5000, 10000, math.Inf(1)},
			outputTax: 0,
			error: false,
		},
		{
			inputAmount: 6001,
			percentages: []float64{0.0, 0.10, 0.20, 0.30},
			amounts: []float64{1000, 5000, 10000, math.Inf(1)},
			outputTax: 600.2,
			error: false,
		},
		{
			inputAmount: 0,
			percentages: []float64{0.0, 0.10, 0.20, 0.30},
			amounts: []float64{1000, 5000, 10000, math.Inf(1)},
			outputTax: 0,
			error: false,
		},
	}
}
