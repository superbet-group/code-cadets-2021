package fizzbuzz_test

type testCase struct {
	inputStart int
	inputEnd   int

	expectedOutput []string
	expectingError bool
}

func getTestCases() []testCase {
	return []testCase{
		{
			inputStart:   1,
			inputEnd:     10,

			expectedOutput: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz"},
			expectingError: false,
		},
		{
			inputStart:   10,
			inputEnd:     20,

			expectedOutput: []string{"Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16", "17", "Fizz", "19", "Buzz"},
			expectingError: false,
		},
		{
			inputStart:   0,
			inputEnd:     10,

			expectedOutput: nil,
			expectingError: true,
		},
		{
			inputStart:   15,
			inputEnd:     15,

			expectedOutput: []string{"FizzBuzz"},
			expectingError: false,
		},
		{
			inputStart:   2,
			inputEnd:     1,

			expectedOutput: nil,
			expectingError: true,
		},
		{
			inputStart:   1,
			inputEnd:     1,

			expectedOutput: []string{"1"},
			expectingError: false,
		},
		{
			inputStart:   3,
			inputEnd:     3,

			expectedOutput: []string{"Fizz"},
			expectingError: false,
		},
		{
			inputStart:   5,
			inputEnd:     5,

			expectedOutput: []string{"Buzz"},
			expectingError: false,
		},
	}
}
