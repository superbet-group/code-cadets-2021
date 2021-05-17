package http

import (
	"code-cadets-2021/homework_2/offerfeed/internal/domain/models"
	"context"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const axilisFeedHomeworkURL = "http://18.193.121.232/axilis-feed-2"

type HomeworkOfferFeed struct {
	httpClient http.Client
	updates    chan models.Odd
}

func NewHomeworkOfferFeed(
	httpClient http.Client,
	updates chan models.Odd,
) *HomeworkOfferFeed {
	return &HomeworkOfferFeed{
		httpClient: httpClient,
		updates:    updates,
	}
}

// close makes sure to close the updates channel only if it isn't already closed
func (h *HomeworkOfferFeed) close() {
	open := true
	select {
	case _, open = <-h.updates:
	default:
	}

	if open {
		close(h.updates)
	}
}

// Start reads Get http response from axilisFeedHomeworkURL and sends it to updates channel
func (h *HomeworkOfferFeed) Start(ctx context.Context) error {
	defer h.close()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(time.Second):
			response, err := h.httpClient.Get(axilisFeedHomeworkURL)
			if err != nil {
				return err
			}

			bodyContent, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return err
			}

			rows := strings.Split(string(bodyContent), "\n")
			for _, row := range rows {
				odd, err := parseOdd(row)
				if err != nil {
					return err
				}

				select {
				case <-ctx.Done():
					return nil
				case h.updates <- odd:
				}
			}
		}
	}
}

func parseOdd(row string) (models.Odd, error) {
	fields := strings.Split(row, ",")

	id := fields[0][2:]
	name := fields[1]
	match := fields[2]
	timestamp := time.Now()
	coefficient, err := strconv.ParseFloat(fields[3], 64)
	if err != nil {
		return models.Odd{}, err
	}

	odd := models.Odd{
		Id:          id,
		Name:        name,
		Match:       match,
		Coefficient: coefficient,
		Timestamp:   timestamp,
	}
	return odd, nil
}
