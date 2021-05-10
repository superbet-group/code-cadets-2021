package main

import (
	"flag"
	"log"

	"github.com/pkg/errors"

	"code-cadets-2021/homework_1/Task-01/fizzbuzz"
)

func main() {
	start := flag.Int("start", 10, "Value (inclusive) from which fizzbuzz starts counting")
	end := flag.Int("end", 20, "Value (inclusive) to which fizzbuzz counts to")
	flag.Parse()

	str, err := fizzbuzz.GetFizzBuzz(*start, *end)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "Error in Fizz Buzz"),
		)
	}

	log.Printf("%s", str)
}
