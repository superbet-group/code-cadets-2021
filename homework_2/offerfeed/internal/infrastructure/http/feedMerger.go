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

func (a *FeedMerger) Start(ctx context.Context) error {
	defer close(a.updates)
	defer log.Printf("shutting down %s", a)

	for _, feed := range a.feeds {
		feed.SetUpdates(a.updates)
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		}
	}
}

func (a *FeedMerger) GetUpdates() chan models.Odd {
	return a.updates
}

func (a *FeedMerger) String() string {
	return "feed merger"
}

type Feed interface {
	SetUpdates(chan models.Odd)
}
