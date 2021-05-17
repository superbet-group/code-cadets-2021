package bootstrap

import (
	"code-cadets-2021/homework_2/offerfeed/internal/domain/models"
	"code-cadets-2021/homework_2/offerfeed/internal/domain/services"
)

func NewProcessingService(feed chan models.Odd, queue chan models.Odd) *services.FeedProcessorService {
	return services.NewFeedProcessorService(feed, queue)
}
