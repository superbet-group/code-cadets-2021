package bootstrap

import (
	stdhttp "net/http"
	"time"

	"code-cadets-2021/homework_2/offerfeed/internal/infrastructure/http"
)

func AxilisOfferFeed() *http.AxilisOfferFeed {
	httpClient := &stdhttp.Client{Timeout: time.Second * 10}
	return http.NewAxilisOfferFeed(httpClient)
}
