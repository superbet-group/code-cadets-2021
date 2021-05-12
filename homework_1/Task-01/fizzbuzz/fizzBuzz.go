package fizzbuzz

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// GetFizzBuzz returns a string which counts from given start (inclusive) to given end (inclusive) according to FizzBuzz rules.
// Returns an error if given start is less than 1 or given start is greater than given end.
func GetFizzBuzz(start, end int) (string, error) {
	if start < 1 {
		return "", errors.New("range start is less than 1")
	}
	if start > end {
		return "", errors.New("range start is greater than range end")
	}

	steps := make([]string, 0)

	for i := start; i <= end; i++ {
		step := ""
		if i%3 == 0 {
			step += "Fizz"
		}

		if i%5 == 0 {
			step += "Buzz"
		}

		if len(step) == 0 {
			step += strconv.Itoa(i)
		}

		steps = append(steps, step)
	}

	return strings.Join(steps, " "), nil
}
