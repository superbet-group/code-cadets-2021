package bootstrap

import (
	"code-cadets-2021/lecture_2/06_offerfeed/internal/infrastructure/http"
)

func FeedComponent(feeds []http.OfferFeed) (*http.FeedComponent, error) {
	return http.NewFeedComponent(feeds)
}
