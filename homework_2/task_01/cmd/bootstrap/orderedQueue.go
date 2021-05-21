package bootstrap

import "code-cadets-2021/homework_2/task_01/internal/infrastructure/queue"

func NewOrderedQueue() *queue.OrderedQueue {
	return queue.NewOrderedQueue()
}
