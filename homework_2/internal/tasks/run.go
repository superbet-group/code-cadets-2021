package tasks

import (
	"context"
	"log"
	"sync"
	"time"
)

func RunTasks(tasks ...Task) {
	wg := &sync.WaitGroup{}
	wg.Add(len(tasks))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	for _, task := range tasks {
		go func(localTask Task) {
			defer cancel()
			defer wg.Done()

			err := localTask.Start(ctx)
			log.Printf(`finished a task with "%v" error`, err)
		}(task)
	}

	log.Print("all tasks running, waiting")
	log.Print("- - - - - - - - - - - - - -")
	wg.Wait()
	log.Print("all tasks finished")
}

type Task interface {
	Start(ctx context.Context) error
}
