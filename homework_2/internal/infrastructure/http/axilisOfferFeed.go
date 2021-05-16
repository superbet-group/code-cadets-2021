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
) *AxilisOfferFeed {
	return &AxilisOfferFeed{
		httpClient: httpClient,
		updates:    make(chan models.Odd),
	}
}

func (a *AxilisOfferFeed) Start(ctx context.Context) error {
	defer close(a.updates)
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

			var decodedAxilisOdds []axilisOfferOdd
			err = json.Unmarshal(bodyContent, &decodedAxilisOdds)
			if err != nil {
				return err
			}

			for _, axilisOdd := range decodedAxilisOdds {
				reMappedOdd := models.Odd{
					Id:          axilisOdd.Id,
					Name:        axilisOdd.Name,
					Match:       axilisOdd.Match,
					Coefficient: axilisOdd.Details.Price,
				}

				select {
				case <-ctx.Done():
					return nil
				case a.updates <- reMappedOdd:
				}
			}
		}
	}
}

func (a *AxilisOfferFeed) GetUpdates() chan models.Odd {
	return a.updates
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
