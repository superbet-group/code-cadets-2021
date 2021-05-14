package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	size    = 1000
	threads = 5
)

func main() {
	c1 := initChannel()
	copy1(c1)

	c2 := initChannel()
	copy2(c2)
}

func initChannel() chan int {
	c := make(chan int, size)
	defer close(c)

	for i := 0; i < size; i++ {
		c <- i
	}
	return c
}

func copy1(c chan int) {
	t := time.Now()

	c2 := make(chan int, size)
	for x := range c {
		// fake http call
		time.Sleep(time.Millisecond)
		c2 <- x
	}

	fmt.Println("copy1 duration", time.Since(t))
}

func copy2(c chan int) {
	t := time.Now()

	wg := &sync.WaitGroup{}
	wg.Add(threads)

	c2 := make(chan int, size)

	for i := 0; i < threads; i++ {
		go func() {
			defer wg.Done()

			for x := range c {
				// fake http call
				time.Sleep(time.Millisecond)
				c2 <- x
			}
		}()
	}

	wg.Wait()

	fmt.Println("copy2 duration", time.Since(t))
}
