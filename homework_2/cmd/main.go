package main

import (
	"fmt"

	"code-cadets-2021/homework_2/offerfeed/cmd/bootstrap"
	"code-cadets-2021/homework_2/offerfeed/internal/domain/models"
	"code-cadets-2021/homework_2/offerfeed/internal/tasks"
)

func main() {
	updatesChannel := make(chan models.Odd)

	offerFeed := bootstrap.NewAxilisOfferFeed(updatesChannel)
	homeworkFeed := bootstrap.NewHomeworkOfferFeed(updatesChannel)

	queue := bootstrap.NewOrderedQueue()
	service := bootstrap.NewProcessingService(offerFeed, queue)
	signalHandler := bootstrap.NewSignalHandler()

	tasks.RunTasks(signalHandler, offerFeed, homeworkFeed, service, queue)

	fmt.Println("program finished gracefully")
}
