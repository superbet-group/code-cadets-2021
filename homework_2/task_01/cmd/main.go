package main

import (
	bootstrap "code-cadets-2021/homework_2/task_01/cmd/bootstrap"
	"code-cadets-2021/homework_2/task_01/internal/tasks"
)

func main() {
	signalHandler := bootstrap.NewSignalHandler()

	feed := bootstrap.NewAxilisOfferFeed()
	queue := bootstrap.NewOrderedQueue()
	processingService := bootstrap.FeedProcessingService(feed, queue)

	// blocking call, start "the application"
	tasks.RunTasks(signalHandler, feed, queue, processingService)

}
