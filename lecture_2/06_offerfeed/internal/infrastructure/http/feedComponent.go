package http

import (
	"context"
	"log"

	"github.com/pkg/errors"

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

// validateOfferFeeds validates that all OfferFeeds have the same updates channel
func validateOfferFeeds(feeds []OfferFeed) (chan models.Odd, error) {

	updates := feeds[0].GetUpdates()
	for _, feed := range feeds {
		if updates != feed.GetUpdates() {
			return nil, errors.New("offer feeds do not contain the same channel")
		}
	}

	return updates, nil
}

func NewFeedComponent(
	feeds []OfferFeed,
) (*FeedComponent, error) {
	updates, err := validateOfferFeeds(feeds)
	if err != nil {
		return nil, err
	}

	return &FeedComponent{
		feeds:   feeds,
		updates: updates,
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
