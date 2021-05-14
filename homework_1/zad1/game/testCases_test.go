package game_test

type testCase struct {
	start int
	end int

	expectedOutput []string
	expectingError bool
}

func getTestCases() []testCase {
	return []testCase{
		{
			start: 1,
			end: 10,

			expectedOutput: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz"},
			expectingError: false,
		},
		{
			start: 10,
			end: 20,

			expectedOutput: []string{"Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16", "17", "Fizz", "19", "Buzz"},
			expectingError: false,
		},
		{
			start: -1,
			end: 20,

			expectedOutput: nil,
			expectingError: true,
		},
		{
			start: 120,
			end: 100,

			expectedOutput: nil,
			expectingError: true,
		},
	}
}