package main

import (
	"fmt"

	"code-cadets-2021/lecture_2/offerfeed/cmd/bootstrap"
	"code-cadets-2021/lecture_2/offerfeed/internal/tasks"
)

func main() {
	signalHandler := bootstrap.SignalHandler()

	feed := bootstrap.AxilisOfferFeed()
	queue := bootstrap.OrderedQueue()
	processingService := bootstrap.FeedProcessingService(feed, queue)

	// blocking call, start "the application"
	tasks.RunTasks(signalHandler, feed, queue, processingService)

	fmt.Println("program finished gracefully")
}
