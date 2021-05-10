package fizzbuzz_test

type testCase struct {
	inputStart int
	inputEnd   int

	expectedOutput string
	expectingError bool
}

func getTestCases() []testCase {
	return []testCase{
		{
			inputStart: 1,
			inputEnd:   10,

			expectedOutput: "1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz",
			expectingError: false,
		},
		{
			inputStart: 5,
			inputEnd:   8,

			expectedOutput: "Buzz Fizz 7 8",
			expectingError: false,
		},
		{
			inputStart: 5,
			inputEnd:   12,

			expectedOutput: "Buzz Fizz 7 8 Fizz Buzz 11 Fizz",
			expectingError: false,
		},
		{
			inputStart: 5,
			inputEnd:   2,

			expectingError: true,
		},
		{
			inputStart: 5,
			inputEnd:   10,

			expectedOutput: "Buzz Fizz 7 8 Fizz Buzz",
			expectingError: false,
		},
		{
			inputStart: -12,
			inputEnd:   10,

			expectingError: true,
		},
		{
			inputStart: 0,
			inputEnd:   10,

			expectingError: true,
		},
		{
			inputStart: 10,
			inputEnd:   15,

			expectedOutput: "Buzz 11 Fizz 13 14 FizzBuzz",
			expectingError: false,
		},
		{
			inputStart: 10,
			inputEnd:   2,

			expectingError: true,
		},
		{
			inputStart: 10,
			inputEnd:   10,

			expectedOutput: "Buzz",
			expectingError: false,
		},
		{
			inputStart: 10,
			inputEnd:   9,

			expectingError: true,
		},
		{
			inputStart: 1,
			inputEnd:   36,

			expectedOutput: "1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz 16 17 Fizz 19 Buzz Fizz 22 23 Fizz Buzz 26 Fizz 28 29 FizzBuzz 31 32 Fizz 34 Buzz Fizz",
			expectingError: false,
		},
		{
			inputStart: 10,
			inputEnd:   20,

			expectedOutput: "Buzz 11 Fizz 13 14 FizzBuzz 16 17 Fizz 19 Buzz",
			expectingError: false,
		},
	}
}
