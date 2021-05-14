package main

import (
	"code-cadets-2021/homework_1/task_03/pokemon"
	"fmt"
	"log"
	"os"
	"strings"
)

func main()  {
	if len(os.Args) != 2 {
		log.Fatalln("wrong input: needed to be name or number of Pokemon")
	}

	input := os.Args[1]
	input = strings.ToLower(input)

	pokemon, err := pokemon.FindLocations(input)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(pokemon))
}