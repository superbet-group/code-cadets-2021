package bootstrap

import (
	stdHttp "net/http"

	"code-cadets-2021/lecture_2/06_offerfeed/internal/domain/models"
	"code-cadets-2021/lecture_2/06_offerfeed/internal/infrastructure/http/offerFeeds"
)

func HomeworkOfferFeed(updates chan models.Odd) *offerFeeds.HomeworkOfferFeed {
	return offerFeeds.NewHomeworkOfferFeed(&stdHttp.Client{}, updates)
}
