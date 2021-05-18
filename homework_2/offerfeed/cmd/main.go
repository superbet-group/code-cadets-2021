package main

import (
	"fmt"

	"code-cadets-2021/homework_2/offerfeed/cmd/bootstrap"
	"code-cadets-2021/homework_2/offerfeed/internal/tasks"
)

func main() {
	signalHandler := bootstrap.SignalHandler()

	axilisOfferFeed := bootstrap.AxilisOfferFeed()
	notAJsonFeed := bootstrap.NotAJsonFeed()

	feedMerger := bootstrap.FeedMerger(axilisOfferFeed, notAJsonFeed)

	queue := bootstrap.OrderedQueue()

	processingService := bootstrap.FeedProcessingService(feedMerger, queue)

	tasks.RunTasks(signalHandler, feedMerger, axilisOfferFeed, notAJsonFeed, queue, processingService)

	fmt.Println("program finished gracefully")
}
