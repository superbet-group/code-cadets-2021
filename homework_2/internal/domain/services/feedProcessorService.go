package services

import (
	"context"

	"code-cadets-2021/homework_2/offerfeed/internal/domain/models"
)

// FeedProcessorService gets values from feedChannel, processes them, and finally sends them to queueChannel
type FeedProcessorService struct {
	feedChannel  chan models.Odd
	queueChannel chan models.Odd
}

func NewFeedProcessorService(feed chan models.Odd, queue chan models.Odd) *FeedProcessorService {
	return &FeedProcessorService{
		feedChannel:  feed,
		queueChannel: queue,
	}
}

// Start reads elements from input channel, multiplies the coefficient by 2, and sends it to output channel
func (f *FeedProcessorService) Start(ctx context.Context) error {
	feedChannel := f.feedChannel
	queueChannel := f.queueChannel
	defer close(queueChannel)

	for odd := range feedChannel {
		odd.Coefficient *= 2
		queueChannel <- odd
	}

	return nil
}
