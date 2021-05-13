package fizzbuzz

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// delimiter between FizzBuzz steps
const delimiter = " "

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

		if i%3 == 0 && i%5 == 0 {
			steps = append(steps, "FizzBuzz")
		} else if i%3 == 0 {
			steps = append(steps, "Fizz")
		} else if i%5 == 0 {
			steps = append(steps, "Buzz")
		} else {
			steps = append(steps, strconv.Itoa(i))
		}
	}

	return strings.Join(steps, delimiter), nil
}
