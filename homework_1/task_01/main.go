package main

import (
"flag"
"fmt"
"log"
"strings"

"code-cadets-2021/homework_1/task_01/fizzBuzz"
)

func parseArgs(start, end int){
	flag.IntVar(&start, "start", 0, "Start value")
	flag.IntVar(&end, "end", 0, "End value")

	flag.Parse()
}

func main() {
	var start, end int

	parseArgs(start, end)

	result, err := fizzBuzz.FizzBuzzPlay(start, end)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", strings.Join(result, " "))
}