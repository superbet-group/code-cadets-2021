package main

import (
	"fmt"

	"code-cadets-2021/homework_2/offerfeed/cmd/bootstrap"
	"code-cadets-2021/homework_2/offerfeed/internal/tasks"
)

func main() {
	offerFeed := bootstrap.NewAxilisOfferFeed()
	queue := bootstrap.NewOrderedQueue()
	service := bootstrap.NewProcessingService(offerFeed, queue)
	signalHandler := bootstrap.NewSignalHandler()

	tasks.RunTasks(signalHandler, offerFeed, service, queue)

	fmt.Println("program finished gracefully")
}
