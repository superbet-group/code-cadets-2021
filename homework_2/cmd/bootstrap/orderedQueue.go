package bootstrap

import (
	"code-cadets-2021/homework_2/offerfeed/internal/domain/models"
	"code-cadets-2021/homework_2/offerfeed/internal/infrastructure/queue"
)

func NewOrderedQueue(source chan models.Odd) *queue.OrderedQueue {
	return queue.NewOrderedQueue(source)
}
