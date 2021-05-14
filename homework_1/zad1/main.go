package main

import (
	"flag"
	"log"

	"github.com/MislavPeric/code-cadets-2021/homework_1/zad1/game"
)

func main() {
	var start, end int

	flag.IntVar(&start, "start", 1, "Value from which the game Fizz Buzz starts")
	flag.IntVar(&end, "end", 100, "Value at which the game Fizz Buzz ends")

	flag.Parse()

	output, err := game.GameLogic(start, end)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v", output)

}