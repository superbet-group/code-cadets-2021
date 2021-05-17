package taxes_test

import (
	"code-cadets-2021/homework_1/task_02/taxes"
	"math"
)

type testCase struct {
	inputAmount float64
	class       []taxes.TaxClass
	outputTax   float64
	error       bool
}

func getTestCases() []testCase {
	return []testCase{
		{
			inputAmount: 7000,
			class: []taxes.TaxClass{
				{1000, 0},
				{5000, 0.10},
				{10000, 0.20},
				{math.Inf(1), 0.30},
			},
			outputTax: 800,
			error:     false,
		},
		{
			inputAmount: 456456,
			class: []taxes.TaxClass{
				{1000, 0},
				{5000, 0.10},
				{10000, 0.20},
				{math.Inf(1), 0.30},
			},
			outputTax: 135336.8,
			error:     false,
		},
		{
			inputAmount: -2500,
			class: []taxes.TaxClass{
				{1000, 0},
				{5000, 0.10},
				{10000, 0.20},
				{math.Inf(1), 0.30},
			},
			error: true,
		},
		{
			inputAmount: 123123,
			class: []taxes.TaxClass{
				{1000, 0},
				{5000, 0.10},
				{10000, 0.20},
				{math.Inf(1), 0.30},
			},
			outputTax: 35336.9,
			error:     false,
		},
		{
			inputAmount: 1000,
			class: []taxes.TaxClass{
				{1000, 0},
				{5000, 0.10},
				{10000, 0.20},
				{math.Inf(1), 0.30},
			},
			outputTax: 0,
			error:     false,
		},
		{
			inputAmount: 6001,
			class: []taxes.TaxClass{
				{1000, 0},
				{5000, 0.10},
				{10000, 0.20},
				{math.Inf(1), 0.30},
			},
			outputTax: 600.2,
			error:     false,
		},
		{
			inputAmount: 0,
			class: []taxes.TaxClass{
				{1000, 0},
				{5000, 0.10},
				{10000, 0.20},
				{math.Inf(1), 0.30},
			},
			outputTax: 0,
			error:     false,
		},
	}
}
