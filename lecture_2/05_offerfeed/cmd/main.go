package main

import (
	"context"
	"fmt"
	"time"

	"code-cadets-2021/lecture_2/05_offerfeed/cmd/bootstrap"
)

func main() {
	offerFeed := bootstrap.NewAxilisOfferFeed()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	go func() {
		err := offerFeed.Start(ctx)
		if err != nil {
			fmt.Println("There was an error.")
		}
	}()

	channel := offerFeed.GetUpdates()
	for upd := range channel {
		fmt.Println(upd)
	}

	cancel()
}
