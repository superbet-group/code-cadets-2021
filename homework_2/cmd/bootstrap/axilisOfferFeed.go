package bootstrap

import (
	stdHttp "net/http"

	"code-cadets-2021/homework_2/offerfeed/internal/infrastructure/http"
)

func NewAxilisOfferFeed() *http.AxilisOfferFeed {
	return http.NewAxilisOfferFeed(stdHttp.Client{})
}
