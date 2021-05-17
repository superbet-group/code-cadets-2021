package bootstrap

import (
	stdHttp "net/http"

	"code-cadets-2021/homework_2/offerfeed/internal/domain/models"
	"code-cadets-2021/homework_2/offerfeed/internal/infrastructure/http/offerFeeds"
)

func NewAxilisOfferFeed(updates chan models.Odd) *offerFeeds.AxilisOfferFeed {
	return offerFeeds.NewAxilisOfferFeed(stdHttp.Client{}, updates)
}
