package main

import (
	"fmt"
	"time"

	"code-cadets-2021/homework_2/offerfeed/cmd/bootstrap"
	"code-cadets-2021/homework_2/offerfeed/internal/tasks"
)

func main() {
	signalHandler := bootstrap.SignalHandler()

	httpClient := bootstrap.HttpClient(time.Second * 10)

	axilisOfferFeed := bootstrap.AxilisOfferFeed(httpClient)
	notAJsonFeed := bootstrap.AnotherAxilisOfferFeed(httpClient)

	feedMerger := bootstrap.FeedMerger(axilisOfferFeed, notAJsonFeed)

	queue := bootstrap.OrderedQueue()

	processingService := bootstrap.FeedProcessingService(feedMerger, queue)

	tasks.RunTasks(signalHandler, axilisOfferFeed, notAJsonFeed, feedMerger, queue, processingService)

	fmt.Println("program finished gracefully")
}
