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

func (a *HomeworkOfferFeed) Start(ctx context.Context) error {
	defer func() {
		ok := true
		select {
		case _, ok = <-a.updates:
		default:
		}

		if ok {
			close(a.updates)
		}
	}()
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(time.Second):
			response, err := a.httpClient.Get(axilisFeedHomeworkURL)
			if err != nil {
				return err
			}

			bodyContent, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return err
			}

			rows := strings.Split(string(bodyContent), "\n")
			for _, row := range rows {
				fields := strings.Split(row, ",")

				id := fields[0][2:]
				name := fields[1]
				match := fields[2]
				timestamp := time.Now()
				coeff, err := strconv.ParseFloat(fields[3], 64)
				if err != nil {
					return err
				}

				odd := models.Odd{
					Id:          id,
					Name:        name,
					Match:       match,
					Coefficient: coeff,
					Timestamp:   timestamp,
				}

				select {
				case <-ctx.Done():
					return nil
				case a.updates <- odd:
				}
			}
		}
	}
}

func (a *HomeworkOfferFeed) GetUpdates() chan models.Odd {
	return a.updates
}
