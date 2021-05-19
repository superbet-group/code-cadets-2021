package bootstrap

import (
	stdHttp "net/http"

	"code-cadets-2021/lecture_2/06_offerfeed/internal/domain/models"
	"code-cadets-2021/lecture_2/06_offerfeed/internal/infrastructure/http/offerFeeds"
)

func AxilisOfferFeed(updates chan models.Odd) *offerFeeds.AxilisOfferFeed {
	return offerFeeds.NewAxilisOfferFeed(&stdHttp.Client{}, updates)
}
