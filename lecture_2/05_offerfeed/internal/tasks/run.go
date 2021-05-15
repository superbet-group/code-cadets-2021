package tasks

import (
	"context"
	"fmt"
	"sync"
)

func RunTasks(tasks ...Task) {
	// run each task in separate goroutine
	// wait for all tasks to finish
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := &sync.WaitGroup{}
	wg.Add(len(tasks))

	for _, task := range tasks {
		go func(task Task) {
			defer wg.Done()
			defer cancel()

			err := task.Start(ctx)
			if err != nil {
				fmt.Println("task error")
			}
		}(task)
	}

	wg.Wait()

	// when first task finishes, signal to the other goroutines that application should stop
}

type Task interface {
	Start(ctx context.Context) error
}
