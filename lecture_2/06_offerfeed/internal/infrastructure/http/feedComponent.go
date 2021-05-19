package http

import (
	"context"
	"log"

	"code-cadets-2021/lecture_2/06_offerfeed/internal/domain/models"
)

type FeedComponent struct {
	feeds   []OfferFeed
	updates chan models.Odd
}

type OfferFeed interface {
	// Start begins feeding offers to updates channel of this OfferFeed
	Start() error
	// GetUpdates returns updates channel of this OfferFeed
	GetUpdates() chan models.Odd
}

func NewFeedComponent(
	feeds []OfferFeed,
) (*FeedComponent, error) {
	return &FeedComponent{
		feeds:   feeds,
		updates: feeds[0].GetUpdates(),
	}, nil
}

func (o *FeedComponent) Start(ctx context.Context) error {
	defer close(o.updates)

	feeds := o.feeds

	for _, feed := range feeds {
		go func(localFeed OfferFeed) {
			err := localFeed.Start()
			log.Printf(`"finished an offer feed with "%v" error`, err)
		}(feed)
	}

	select {
	case <-ctx.Done():
	}

	return nil
}

func (o *FeedComponent) GetUpdates() chan models.Odd {
	return o.updates
}
