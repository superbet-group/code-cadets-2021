package bootstrap

import (
	stdHttp "net/http"

	"code-cadets-2021/homework_2/offerfeed/internal/domain/models"
	"code-cadets-2021/homework_2/offerfeed/internal/infrastructure/http/offerFeeds"
)

func NewHomeworkOfferFeed(updates chan models.Odd) *offerFeeds.HomeworkOfferFeed {
	return offerFeeds.NewHomeworkOfferFeed(stdHttp.Client{}, updates)
}
