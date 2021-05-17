package bootstrap

import (
	"code-cadets-2021/homework_2/offerfeed/internal/infrastructure/http"
)

func NewFeedComponent(feeds []http.OfferFeed) (*http.FeedComponent, error) {
	return http.NewFeedComponent(feeds)
}
