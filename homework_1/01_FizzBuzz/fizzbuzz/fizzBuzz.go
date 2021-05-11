package fizzbuzz

import (
	"strconv"

	"github.com/pkg/errors"
)

func GetFizzBuzzSolution(start int, end int) ([]string, error) {
	if start > end {
		return nil, errors.New("Start is greater than end.")
	}

	var solution []string

	for n := start; n <= end; n++ {
		if n % 3 == 0 {
			if n % 5 == 0 {
				solution = append(solution, "FizzBuzz")
			} else {
				solution = append(solution, "Fizz")
			}
		} else if n % 5 == 0 {
			solution = append(solution, "Buzz")
		} else {
			solution = append(solution, strconv.Itoa(n))
		}
	}

	return solution, nil
}