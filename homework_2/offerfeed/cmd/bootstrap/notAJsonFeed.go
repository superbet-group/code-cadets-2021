package bootstrap

import (
	stdhttp "net/http"

	"code-cadets-2021/homework_2/offerfeed/internal/infrastructure/http"
)

func NotAJsonFeed() *http.NotAJsonFeed {
	httpClient := &stdhttp.Client{}
	return http.NewNotAJsonFeed(httpClient)
}

