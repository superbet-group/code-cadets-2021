package main

import (
	"log"

	"code-cadets-2021/lecture_2/01_interfaces/stacklibfer"
)

func main() {
	s := stacklibfer.New()

	pushing(s)
}

func pushing(pushPopper Pusher) {
	pushPopper.Push(1)
	pushPopper.Push(2)
	pushPopper.Push(3)
	pushPopper.Push(4)

	log.Println(pushPopper)
}

type Pusher interface {
	Push(a int)
}
