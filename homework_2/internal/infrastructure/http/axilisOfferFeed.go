package http

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"code-cadets-2021/homework_2/offerfeed/internal/domain/models"
)

const axilisFeedURL = "http://18.193.121.232/axilis-feed"

type AxilisOfferFeed struct {
	httpClient http.Client
	updates    chan models.Odd
}

func NewAxilisOfferFeed(
	httpClient http.Client,
	updates chan models.Odd,
) *AxilisOfferFeed {
	return &AxilisOfferFeed{
		httpClient: httpClient,
		updates:    updates,
	}
}

func (a *AxilisOfferFeed) close() {
	open := true
	select {
	case _, open = <-a.updates:
	default:
	}

	if open {
		close(a.updates)
	}
}

func (a *AxilisOfferFeed) Start(ctx context.Context) error {
	defer a.close()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(time.Second):
			response, err := a.httpClient.Get(axilisFeedURL)
			if err != nil {
				return err
			}

			bodyContent, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return err
			}

			var decodedOfferOdds []axilisOfferOdd
			err = json.Unmarshal(bodyContent, &decodedOfferOdds)
			if err != nil {
				return err
			}

			for _, odd := range decodedOfferOdds {
				mappedOdd := models.Odd{
					Id:          odd.Id,
					Name:        odd.Name,
					Match:       odd.Match,
					Coefficient: odd.Details.Price,
					Timestamp:   time.Now(),
				}

				select {
				case <-ctx.Done():
					return nil
				case a.updates <- mappedOdd:
				}
			}
		}
	}
}

type axilisOfferOdd struct {
	Id      string
	Name    string
	Match   string
	Details axilisOfferOddDetails
}

type axilisOfferOddDetails struct {
	Price float64
}
