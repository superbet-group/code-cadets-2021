package fizzBuzz_test

type testCase struct {
	inputStart int
	inputEnd   int
	expectedOutput []string
	expectingError bool
}

func getTestCases() []testCase {
	return []testCase{
		{
			inputStart: 1,
			inputEnd:   10,

			expectedOutput: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz"},
			expectingError: false,
		},
		{
			inputStart: 10,
			inputEnd:   20,

			expectedOutput: []string{"Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16", "17", "Fizz", "19", "Buzz"},
			expectingError: false,
		},
		{
			inputStart: 1,
			inputEnd:   30,

			expectedOutput: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16", "17", "Fizz", "19", "Buzz", "Fizz", "22", "23", "Fizz", "Buzz", "26", "Fizz", "28", "29", "FizzBuzz"},
			expectingError: false,
		},
		{
			inputStart: 15,
			inputEnd:   25,

			expectedOutput: []string{"FizzBuzz", "16", "17", "Fizz", "19", "Buzz", "Fizz", "22", "23", "Fizz", "Buzz"},
			expectingError: false,
		},
		{
			inputStart: 80,
			inputEnd:   100,

			expectedOutput: []string{"Buzz", "Fizz", "82", "83", "Fizz", "Buzz", "86", "Fizz", "88", "89", "FizzBuzz", "91", "92", "Fizz", "94", "Buzz", "Fizz", "97", "98", "Fizz", "Buzz"},
			expectingError: false,
		},
		{
			inputStart: 10,
			inputEnd:   5,

			expectedOutput: nil,
			expectingError: true,
		},
		{
			inputStart: 0,
			inputEnd:   10,

			expectingError: true,
		},
		{
			inputStart: 12,
			inputEnd:   -4,

			expectingError: true,
		},
		{
			inputStart: 1,
			inputEnd:   0,

			expectingError: true,
		},
		{
			inputStart: 522,
			inputEnd:   -1,

			expectingError: true,
		},
	}
}