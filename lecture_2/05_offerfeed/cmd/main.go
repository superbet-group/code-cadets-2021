package main

import (
	bootstrap "code-cadets-2021/lecture_2/05_offerfeed/cmd/bootstrap"
	"code-cadets-2021/lecture_2/05_offerfeed/internal/tasks"
)

func main() {
	signalHandler := bootstrap.NewSignalHandler()

	feed := bootstrap.NewAxilisOfferFeed()
	queue := bootstrap.NewOrderedQueue()
	processingService := bootstrap.FeedProcessingService(feed, queue)

	// blocking call, start "the application"
	tasks.RunTasks(signalHandler, feed, queue, processingService)

}
