package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"code-cadets-2021/lecture_2/05_offerfeed/cmd/bootstrap"
)

func main() {
	feed := bootstrap.NewAxilisOfferFeed()
	queue := bootstrap.NewOrderedQueue()

	feedProcessor := bootstrap.NewFeedProcessorService(feed, queue)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	wg := &sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		err := feed.Start(ctx)
		if err != nil {
			fmt.Println("feed error")
		}
	}()

	go func() {
		defer wg.Done()
		err := queue.Start(ctx)
		if err != nil {
			fmt.Println("queue error")
		}
	}()

	go func() {
		defer wg.Done()
		err := feedProcessor.Start(ctx)
		if err != nil {
			fmt.Println("processor error")
		}
	}()

	wg.Wait()

	fmt.Println("program finished gracefully")
}
