package main

import (
	"fmt"
	"log"

	"code-cadets-2021/homework_2/offerfeed/cmd/bootstrap"
	"code-cadets-2021/homework_2/offerfeed/internal/domain/models"
	"code-cadets-2021/homework_2/offerfeed/internal/infrastructure/http"
	"code-cadets-2021/homework_2/offerfeed/internal/tasks"
)

func main() {
	updatesChannel := make(chan models.Odd)
	queueChannel := make(chan models.Odd)

	offerFeed := bootstrap.NewAxilisOfferFeed(updatesChannel)
	homeworkFeed := bootstrap.NewHomeworkOfferFeed(updatesChannel)
	feedComponent, err := bootstrap.NewFeedComponent([]http.OfferFeed{offerFeed, homeworkFeed})
	if err != nil {
		log.Fatalln(err, "error while constructing OfferFeedComponent")
	}

	queue := bootstrap.NewOrderedQueue(queueChannel)

	service := bootstrap.NewProcessorService(feedComponent, queue)

	signalHandler := bootstrap.NewSignalHandler()

	tasks.RunTasks(signalHandler, feedComponent, service, queue)
	fmt.Println("program finished gracefully")
}
