package main

import (
	"flag"
	"fmt"
	"strings"

	"code-cadets-2021/homework_1/01_FizzBuzz/fizzbuzz"
)

func main() {
	start := flag.Int("start", 1, "Value (inclusive) from which to start counting.")
	end := flag.Int("end", 10, "Value (inclusive) to count to.")

	flag.Parse()

	solution, err := fizzbuzz.GetFizzBuzzSolution(*start, *end)
	if err == nil {
		fmt.Println(strings.Join(solution, " "))
	}
}
