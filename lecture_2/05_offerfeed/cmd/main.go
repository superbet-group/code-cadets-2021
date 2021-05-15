package main

import (
	"context"
	"fmt"
	"time"

	"code-cadets-2021/lecture_2/05_offerfeed/cmd/bootstrap"
	"code-cadets-2021/lecture_2/05_offerfeed/internal/domain/models"
)

func main() {
	queue := bootstrap.NewOrderedQueue()
	source := queue.GetSource()

	source <- models.Odd{
		Id:          "1",
		Name:        "",
		Match:       "",
		Coefficient: 0,
		Timestamp:   time.Time{},
	}

	close(source)

	err := queue.Start(context.Background())
	if err != nil {
		fmt.Println("there was an error")
	} else {
		fmt.Println("program finished gracefully")
	}
}
