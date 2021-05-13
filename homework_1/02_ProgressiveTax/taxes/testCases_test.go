package taxes_test

import "math"

type newTaxBracketsTestCase struct {
	inputThresholds []float32
	inputTaxRates   []float32

	expectingError bool
}

type calculateProgressiveTaxTestCase struct {
	inputValue      float32
	inputThresholds []float32
	inputTaxRates   []float32

	expectedOutput float32
	expectingError bool
}

func getCreateBracketsTestCases() []newTaxBracketsTestCase {
	return []newTaxBracketsTestCase{
		{
			inputThresholds: []float32{1000, 5000, 10000, math.MaxFloat32},
			inputTaxRates:   []float32{0, 0.1, 0.2, 0.3},

			expectingError: false,
		},
		{
			inputThresholds: []float32{1000, 5000, math.MaxFloat32},
			inputTaxRates:   []float32{0, 0.1, 0.2, 0.3},

			expectingError: true,
		},
		{
			inputThresholds: []float32{1000, 5000, math.MaxFloat32},
			inputTaxRates:   []float32{0, 0.1},

			expectingError: true,
		},
		{
			inputThresholds: []float32{1000, 5000, 10000, math.MaxFloat32},
			inputTaxRates:   []float32{},

			expectingError: true,
		},
		{
			inputThresholds: []float32{},
			inputTaxRates:   []float32{0, 0.1, 0.2, 0.3},

			expectingError: true,
		},
		{
			inputThresholds: []float32{},
			inputTaxRates:   []float32{},

			expectingError: false,
		},
		{
			inputThresholds: []float32{1000, 5000, 10000, math.MaxFloat32},
			inputTaxRates:   []float32{0.9, 0.1, 0.2, 0.3},

			expectingError: false,
		},
		{
			inputThresholds: []float32{1000, 5000, 10000, math.MaxFloat32},
			inputTaxRates:   []float32{0, 0.1, 0.2, -0.3},

			expectingError: true,
		},
	}
}

func getCalculateProgressiveTaxTestCases() []calculateProgressiveTaxTestCase {
	return []calculateProgressiveTaxTestCase{
		{
			inputValue:      7000,
			inputThresholds: []float32{1000, 5000, 10000, math.MaxFloat32},
			inputTaxRates:   []float32{0, 0.1, 0.2, 0.3},

			expectedOutput: 800,
			expectingError: false,
		},
		{
			inputValue:      -7000,
			inputThresholds: []float32{1000, 5000, 10000, math.MaxFloat32},
			inputTaxRates:   []float32{0, 0.1, 0.2, 0.3},

			expectedOutput: 0,
			expectingError: true,
		},
		{
			inputValue:      0,
			inputThresholds: []float32{1000, 5000, 10000, math.MaxFloat32},
			inputTaxRates:   []float32{0, 0.1, 0.2, 0.3},

			expectedOutput: 0,
			expectingError: false,
		},
		{
			inputValue:      20000,
			inputThresholds: []float32{10000, 20000, math.MaxFloat32},
			inputTaxRates:   []float32{0.1, 0.2, 0.3},

			expectedOutput: 3000,
			expectingError: false,
		},
		{
			inputValue:      25000,
			inputThresholds: []float32{10000, 20000, math.MaxFloat32},
			inputTaxRates:   []float32{0.1, 0.2, 0.3},

			expectedOutput: 4500,
			expectingError: false,
		},
		{
			inputValue:      10000,
			inputThresholds: []float32{10000, 20000, math.MaxFloat32},
			inputTaxRates:   []float32{0.1, 0.2, 0.3},

			expectedOutput: 1000,
			expectingError: false,
		},
		{
			inputValue:      10000,
			inputThresholds: []float32{7550, 10000, math.MaxFloat32},
			inputTaxRates:   []float32{0.1, 0.15, 0.2},

			expectedOutput: 1122.5,
			expectingError: false,
		},
	}
}
