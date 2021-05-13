package fizzBuzz

import (
	"strconv"

	"github.com/pkg/errors"
)


func FizzBuzzPlay(start, end int) ([]string, error) {
	if start <= 0 {
		return nil, errors.New("invalid value, range start should be greater than 0")
	}
	if end <= 0 {
		return nil, errors.New("invalid value, range end should be greater than 0")
	}
	if start > end {
		return nil, errors.New("range start is greater than range end")
	}


	var output []string

	for i := start; i <= end; i++ {
		if i % 15 == 0 {
			output = append(output, "FizzBuzz")
		} else if i % 5 == 0 {
			output = append(output, "Buzz")
		} else if i % 3 == 0 {
			output = append(output, "Fizz")
		} else {
			output = append(output, strconv.Itoa(i))
		}
	}
	return output, nil
}