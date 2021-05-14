
package game

import (
	"github.com/pkg/errors"
	"strconv"
)

func GameLogic(start, end int) ([]string, error) {
	if start > end {
		return nil, errors.New("start is greater than end is")
	}

	if start < 0 {
		return nil, errors.New("start can't be less than 0")
	}

	var output []string

	for i := start; i <= end; i++ {
		if i % 3 == 0 && i % 5 == 0 {
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