package services

import (
	"context"

	"code-cadets-2021/lecture_2/offerfeed/internal/domain/models"
)

type FeedProcessorService struct {
	feed  Feed
	queue Queue
}

func NewFeedProcessorService(
	feed Feed,
	queue Queue,
) *FeedProcessorService {
	return &FeedProcessorService{
		feed:  feed,
		queue: queue,
	}
}

func (f *FeedProcessorService) Start(ctx context.Context) error {
	updates := f.feed.GetUpdates()
	source := f.queue.GetSource()

	for update := range updates {
		source <- update
	}

	close(source)

	return nil
}

type Feed interface {
	GetUpdates() <-chan models.Odd
}

type Queue interface {
	GetSource() chan<- models.Odd
}
