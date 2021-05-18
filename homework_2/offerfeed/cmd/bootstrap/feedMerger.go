package bootstrap

import (
	"code-cadets-2021/homework_2/offerfeed/internal/infrastructure/http"
)

func FeedMerger(feeds ...http.Feed) *http.FeedMerger {
	return http.NewFeedMerger(feeds...)
}
