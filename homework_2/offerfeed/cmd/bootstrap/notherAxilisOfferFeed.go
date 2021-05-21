package bootstrap

import (
	stdhttp "net/http"

	"code-cadets-2021/homework_2/offerfeed/internal/infrastructure/http"
)

func AnotherAxilisOfferFeed(httpClient *stdhttp.Client) *http.AnotherAxilisOfferFeed {
	return http.NewAnotherAxilisOfferFeed(httpClient)
}
