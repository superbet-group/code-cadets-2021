package bootstrap

import (
	stdhttp "net/http"
	"time"

	"code-cadets-2021/homework_2/offerfeed/internal/infrastructure/http"
)

func NotAJsonFeed() *http.NotAJsonFeed {
	httpClient := &stdhttp.Client{Timeout: time.Second * 10}
	return http.NewNotAJsonFeed(httpClient)
}

