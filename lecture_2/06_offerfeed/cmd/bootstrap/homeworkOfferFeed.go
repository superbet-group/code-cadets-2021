package bootstrap

import (
	stdhttp "net/http"

	"code-cadets-2021/lecture_2/06_offerfeed/internal/domain/models"
	"code-cadets-2021/lecture_2/06_offerfeed/internal/infrastructure/http/offerFeeds"
)

func HomeworkOfferFeed(httpClient *stdhttp.Client, updates chan models.Odd) *offerFeeds.HomeworkOfferFeed {
	return offerFeeds.NewHomeworkOfferFeed(httpClient, updates)
}
