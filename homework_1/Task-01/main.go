package main

import (
	"log"

	"github.com/pkg/errors"

	"code-cadets-2021/homework_1/Task-01/FizzBuzz"
)

func main() {
	var start, end int

	str, err := FizzBuzz.FizzBuzz(start, end)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "HTTP get towards yesno API"),
		)
	}

	log.Printf("%s", str)
}
