package http

import (
	"context"
	"log"

	"code-cadets-2021/homework_2/offerfeed/internal/domain/models"
)

type FeedMerger struct {
	feeds   []Feed
	updates chan models.Odd
}

func NewFeedMerger(feeds ...Feed) *FeedMerger {
	return &FeedMerger{feeds: feeds, updates: make(chan models.Odd)}
}

func (f *FeedMerger) Start(ctx context.Context) error {
	defer close(f.updates)
	defer log.Printf("shutting down %s", f)

	for _, feed := range f.feeds {
		feed.SetUpdates(f.updates)
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		}
	}
}

func (f *FeedMerger) GetUpdates() chan models.Odd {
	return f.updates
}

func (f *FeedMerger) String() string {
	return "feed merger"
}

type Feed interface {
	SetUpdates(chan models.Odd)
}
