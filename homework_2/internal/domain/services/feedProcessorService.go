package services

import (
	"context"

	"code-cadets-2021/homework_2/offerfeed/internal/domain/models"
)

type FeedProcessorService struct {
	feed  Feed
	queue Queue
}

func NewFeedProcessorService(feed Feed, queue Queue) *FeedProcessorService {
	return &FeedProcessorService{
		feed:  feed,
		queue: queue,
	}
}

func (f *FeedProcessorService) Start(ctx context.Context) error {
	feedChannel := f.feed.GetUpdates()
	queueChannel := f.queue.GetSource()
	defer close(queueChannel)

	for odd := range feedChannel {
		odd.Coefficient *= 2
		queueChannel <- odd
	}

	return nil
}

type Feed interface {
	GetUpdates() chan models.Odd
}

type Queue interface {
	GetSource() chan models.Odd
}
