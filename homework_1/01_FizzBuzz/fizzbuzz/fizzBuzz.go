package fizzbuzz

import (
	"strconv"

	"github.com/pkg/errors"
)

func GetFizzBuzzSolution(start int, end int) ([]string, error) {
	if start > end {
		return nil, errors.New("Start is greater than end.")
	}

	if start < 1 {
		return nil, errors.New("Start is lower than 1.")
	}

	var solution []string

	for ;start <= end; start++ {
		if start % 3 == 0 {
			if start % 5 == 0 {
				solution = append(solution, "FizzBuzz")
			} else {
				solution = append(solution, "Fizz")
			}
		} else if start % 5 == 0 {
			solution = append(solution, "Buzz")
		} else {
			solution = append(solution, strconv.Itoa(start))
		}
	}

	return solution, nil
}