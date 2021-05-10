// Package fizzbuzz offers an implementation of the FizzBuzz game
// more info: https://en.wikipedia.org/wiki/Fizz_buzz
package fizzbuzz

import (
	"strconv"

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

	output := ""

	for i := start; i <= end; i++ {
		str := ""
		if i%3 == 0 {
			str += "Fizz"
		}

		if i%5 == 0 {
			str += "Buzz"
		}

		if len(str) == 0 {
			str += strconv.Itoa(i)
		}

		output += str + " "
	}
	return output[:len(output)-1], nil
}
