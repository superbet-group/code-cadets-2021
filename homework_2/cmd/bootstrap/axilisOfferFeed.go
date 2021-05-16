package bootstrap

import (
	stdHttp "net/http"

	"code-cadets-2021/homework_2/offerfeed/internal/domain/models"
	"code-cadets-2021/homework_2/offerfeed/internal/infrastructure/http"
)

func NewAxilisOfferFeed(updates chan models.Odd) *http.AxilisOfferFeed {
	return http.NewAxilisOfferFeed(stdHttp.Client{}, updates)
}
