package bootstrap

import "code-cadets-2021/homework_2/offerfeed/internal/infrastructure/queue"

func OrderedQueue() *queue.OrderedQueue {
	return queue.NewOrderedQueue()
}
