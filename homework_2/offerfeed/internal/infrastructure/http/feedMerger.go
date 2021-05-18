package http

import (
	"context"
	"log"

	"code-cadets-2021/homework_2/offerfeed/internal/domain/models"
)

type FeedMerger struct {
	Feeds   []Feed
	Updates chan models.Odd
}

func NewFeedMerger(feeds ...Feed) *FeedMerger {
	return &FeedMerger{Feeds: feeds, Updates: make(chan models.Odd)}
}

func (a *FeedMerger) Start(ctx context.Context) error {
	defer close(a.Updates)
	defer log.Printf("shutting down %s", a)

	for _, feed := range a.Feeds {
		feed.SetUpdates(a.Updates)
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		}
	}
}

func (a *FeedMerger) GetUpdates() chan models.Odd {
	return a.Updates
}

func (a *FeedMerger) String() string {
	return "feed merger"
}

type Feed interface {
	SetUpdates(chan models.Odd)
}