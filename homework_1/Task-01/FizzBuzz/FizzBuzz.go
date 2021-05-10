package FizzBuzz

import (
	"github.com/pkg/errors"
)

func FizzBuzz(start, end int) (string, error) {
	if start < 1 {
		return "", errors.New("range start is less than 1")
	}
	if start > end {
		return "", errors.New("range start is greater than range end")
	}

	output := ""

	for i := start; i <= end; i++ {
		if i%3 == 0 {
			output += " Fizz "
		}

		if i%5 == 0 {
			output += " Buzz "
		}

		output += ","
	}
	return output[:len(output)-1], nil
}
