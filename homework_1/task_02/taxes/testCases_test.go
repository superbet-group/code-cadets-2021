package taxes_test

type testCase struct {
	input float64
	output float64
	error bool
}

func getTestCases() []testCase {
	return []testCase {
		{
			input: 7000,
			output: 800,
			error: false,
		},
		{
			input: 456456,
			output: 135336.8,
			error: false,
		},
		{
			input: -2500,
			error: true,
		},
		{
			input: 123123,
			output: 35336.9,
			error: false,
		},
		{
			input: 1000,
			output: 0,
			error: false,
		},
		{
			input: 6001,
			output: 600.2,
			error: false,
		},
		{
			input: 0,
			output: 0,
			error: false,
		},
	}
}