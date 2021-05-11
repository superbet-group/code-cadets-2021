package taxes_test

type newTaxBracketsTestCase struct {
	inputBrackets []float32
	inputTaxes    []float32

	expectingError bool
}

type calculateProgressiveTaxTestCase struct {
	inputValue    float32
	inputBrackets []float32
	inputTaxes    []float32

	expectedOutput float32
	expectingError bool
}

func getNewTaxBracketsTestCases() []newTaxBracketsTestCase {
	return []newTaxBracketsTestCase{
		{
			inputBrackets: []float32{0, 1000, 5000, 10000},
			inputTaxes:    []float32{0, 0.1, 0.2, 0.3},

			expectingError: false,
		},
		{
			inputBrackets: []float32{0, 1000, 5000},
			inputTaxes:    []float32{0, 0.1, 0.2, 0.3},

			expectingError: true,
		},
		{
			inputBrackets: []float32{0, 1000, 5000, 10000},
			inputTaxes:    []float32{0, 0.1, 0.2},

			expectingError: true,
		},
		{
			inputBrackets: []float32{1, 1000, 5000, 10000},
			inputTaxes:    []float32{0, 0.1, 0.2, 0.3},

			expectingError: true,
		},
		{
			inputBrackets: []float32{0, 1000, 5000, 10000},
			inputTaxes:    []float32{},

			expectingError: true,
		},
		{
			inputBrackets: []float32{},
			inputTaxes:    []float32{0, 0.1, 0.2, 0.3},

			expectingError: true,
		},
		{
			inputBrackets: []float32{},
			inputTaxes:    []float32{},

			expectingError: true,
		},
		{
			inputBrackets: []float32{0, 5000, 5000, 10000},
			inputTaxes:    []float32{0, 0.1, 0.2, 0.3},

			expectingError: true,
		},
		{
			inputBrackets: []float32{0, 1100, 1000, 10000},
			inputTaxes:    []float32{0, 0.1, 0.2, 0.3},

			expectingError: true,
		},
		{
			inputBrackets: []float32{0, 1000, 5000, 10000},
			inputTaxes:    []float32{0.9, 0.1, 0.2, 0.3},

			expectingError: false,
		},
		{
			inputBrackets: []float32{0, 1000, 5000, 10000},
			inputTaxes:    []float32{0, 0.1, 0.2, -0.3},

			expectingError: true,
		},
	}
}

func getCalculateProgressiveTaxTestCases() []calculateProgressiveTaxTestCase {
	return []calculateProgressiveTaxTestCase{
		{
			inputValue:    7000,
			inputBrackets: []float32{0, 1000, 5000, 10000},
			inputTaxes:    []float32{0, 0.1, 0.2, 0.3},

			expectedOutput: 800,
			expectingError: false,
		},
		{
			inputValue:    -7000,
			inputBrackets: []float32{0, 1000, 5000, 10000},
			inputTaxes:    []float32{0, 0.1, 0.2, 0.3},

			expectedOutput: 0,
			expectingError: true,
		},
		{
			inputValue:    0,
			inputBrackets: []float32{0, 1000, 5000, 10000},
			inputTaxes:    []float32{0, 0.1, 0.2, 0.3},

			expectedOutput: 0,
			expectingError: false,
		},
		{
			inputValue:    20000,
			inputBrackets: []float32{0, 10000, 20000},
			inputTaxes:    []float32{0.1, 0.2, 0.3},

			expectedOutput: 3000,
			expectingError: false,
		},
		{
			inputValue:    25000,
			inputBrackets: []float32{0, 10000, 20000},
			inputTaxes:    []float32{0.1, 0.2, 0.3},

			expectedOutput: 4500,
			expectingError: false,
		},
		{
			inputValue:    10000,
			inputBrackets: []float32{0, 10000, 20000},
			inputTaxes:    []float32{0.1, 0.2, 0.3},

			expectedOutput: 1000,
			expectingError: false,
		},
		{
			inputValue:    10000,
			inputBrackets: []float32{0, 7550, 10000},
			inputTaxes:    []float32{0.1, 0.15, 0.2},

			expectedOutput: 1122.5,
			expectingError: false,
		},
	}
}
