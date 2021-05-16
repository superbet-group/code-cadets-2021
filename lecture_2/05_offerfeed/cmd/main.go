package main

import (
	"code-cadets-2021/lecture_2/05_offerfeed/internal/domain/services"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"

	"code-cadets-2021/lecture_2/05_offerfeed/cmd/bootstrap"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	offerFeed := bootstrap.NewAxilisOfferFeed()

	go func() {
		err := offerFeed.Start(ctx)
		if err != nil {
			log.Fatal(errors.WithMessage(err, "error starting offer feed"))
		}
	}()

	queue := bootstrap.NewOrderedQueue()

	service := services.NewFeedProcessorService(offerFeed, queue)
	go func() {
		err := service.Start(ctx)
		if err != nil {
			log.Fatal(errors.WithMessage(err, "error starting offer queue"))
		}
	}()

	go func() {
		err := queue.Start(ctx)
		if err != nil {
			log.Fatal(errors.WithMessage(err, "error starting offer queue"))
		}
	}()

	time.Sleep(time.Second * 6)

	fmt.Println("program finished gracefully")
}
