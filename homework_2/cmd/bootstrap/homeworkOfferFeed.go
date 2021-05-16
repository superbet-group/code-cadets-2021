package bootstrap

import (
	stdHttp "net/http"

	"code-cadets-2021/homework_2/offerfeed/internal/domain/models"
	"code-cadets-2021/homework_2/offerfeed/internal/infrastructure/http"
)

func NewHomeworkOfferFeed(updates chan models.Odd) *http.HomeworkOfferFeed {
	return http.NewHomeworkOfferFeed(stdHttp.Client{}, updates)
}
