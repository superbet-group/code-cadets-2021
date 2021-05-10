package fizzbuzz

import (
	"strconv"

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
