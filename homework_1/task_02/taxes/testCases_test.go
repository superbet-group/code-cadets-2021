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
				{UpperThreshold: 1000, Percentage: 0},
				{UpperThreshold: 5000, Percentage: 0.10},
				{UpperThreshold: 10000, Percentage: 0.20},
				{UpperThreshold: math.Inf(1), Percentage: 0.30},
			},
			outputTax: 800,
			error:     false,
		},
		{
			inputAmount: 456456,
			class: []taxes.TaxClass{
				{UpperThreshold: 1000, Percentage: 0},
				{UpperThreshold: 5000, Percentage: 0.10},
				{UpperThreshold: 10000, Percentage: 0.20},
				{UpperThreshold: math.Inf(1), Percentage: 0.30},
			},
			outputTax: 135336.8,
			error:     false,
		},
		{
			inputAmount: -2500,
			class: []taxes.TaxClass{
				{UpperThreshold: 1000, Percentage: 0},
				{UpperThreshold: 5000, Percentage: 0.10},
				{UpperThreshold: 10000, Percentage: 0.20},
				{UpperThreshold: math.Inf(1), Percentage: 0.30},
			},
			error: true,
		},
		{
			inputAmount: 123123,
			class: []taxes.TaxClass{
				{UpperThreshold: 1000, Percentage: 0.0},
				{UpperThreshold: 5000, Percentage: 0.10},
				{UpperThreshold: 10000, Percentage: 0.20},
				{UpperThreshold: math.Inf(1), Percentage: 0.30},
			},
			outputTax: 35336.9,
			error:     false,
		},
		{
			inputAmount: 1000,
			class: []taxes.TaxClass{
				{UpperThreshold: 1000, Percentage: 0},
				{UpperThreshold: 5000, Percentage: 0.10},
				{UpperThreshold: 10000, Percentage: 0.20},
				{UpperThreshold: math.Inf(1), Percentage: 0.30},
			},
			outputTax: 0,
			error:     false,
		},
		{
			inputAmount: 6001,
			class: []taxes.TaxClass{
				{UpperThreshold: 1000, Percentage: 0},
				{UpperThreshold: 5000, Percentage: 0.10},
				{UpperThreshold: 10000, Percentage: 0.20},
				{UpperThreshold: math.Inf(1), Percentage: 0.30},
			},
			outputTax: 600.2,
			error:     false,
		},
		{
			inputAmount: 0,
			class: []taxes.TaxClass{
				{UpperThreshold: 1000, Percentage: 0},
				{UpperThreshold: 5000, Percentage: 0.10},
				{UpperThreshold: 10000, Percentage: 0.20},
				{UpperThreshold: math.Inf(1), Percentage: 0.30},
			},
			outputTax: 0,
			error:     false,
		},
		{
			inputAmount: 10000,
			class: []taxes.TaxClass{
				{UpperThreshold: 1000, Percentage: 0},
				{UpperThreshold: 5000, Percentage: 0.10},
				{UpperThreshold: 1000, Percentage: 0.20},
				{UpperThreshold: math.Inf(1), Percentage: 0.30},
			},
			outputTax: 0,
			error:     true,
		},
		{
			inputAmount: 10000,
			class: []taxes.TaxClass{
				{UpperThreshold: -1000, Percentage: 0},
				{UpperThreshold: 5000, Percentage: 0.10},
				{UpperThreshold: 10000, Percentage: 0.20},
				{UpperThreshold: math.Inf(1), Percentage: 0.30},
			},
			outputTax: 0,
			error:     true,
		},
		{
			inputAmount: 19000,
			class: []taxes.TaxClass{
				{UpperThreshold: 1000, Percentage: 0},
				{UpperThreshold: 7000, Percentage: 0.10},
				{UpperThreshold: 12000, Percentage: 0.25},
				{UpperThreshold: math.Inf(1), Percentage: 0.35},
			},
			outputTax: 4300,
			error:     false,
		},
		{
			inputAmount: 13500,
			class: []taxes.TaxClass{
				{UpperThreshold: 1200, Percentage: 0},
				{UpperThreshold: 4800, Percentage: 0.10},
				{UpperThreshold: 9600, Percentage: 0.25},
				{UpperThreshold: math.Inf(1), Percentage: 0.35},
			},
			outputTax: 2925,
			error:     false,
		},
		{
			inputAmount: 13500,
			class: []taxes.TaxClass{
				{UpperThreshold: 1200, Percentage: 0},
				{UpperThreshold: 4800, Percentage: 0.10},
				{UpperThreshold: 9600, Percentage: 0.25},
				{UpperThreshold: 8240, Percentage: 0.35},
			},
			outputTax: 0,
			error:     true,
		},
	}
}
