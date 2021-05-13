package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"code-cadets-2021/lecture_2/offerfeed/internal/domain/models"
)

const url = "http://18.193.121.232/axilis-feed"

type AxilisOfferFeed struct {
	updates    chan models.Odd
	httpClient *http.Client
}

func NewAxilisOfferFeed(
	httpClient *http.Client,
) *AxilisOfferFeed {
	return &AxilisOfferFeed{
		updates:    make(chan models.Odd),
		httpClient: httpClient,
	}
}

func (a *AxilisOfferFeed) Start(ctx context.Context) error {
	defer close(a.updates)

	for {
		select {
		case <-ctx.Done():
			return nil

		case <-time.After(time.Second * 3):
			response, err := a.httpClient.Get(url)
			if err != nil {
				log.Println("axilis offer feed, http get", err)
				continue
			}
			a.processResponse(response)
		}
	}
}

func (a *AxilisOfferFeed) GetUpdates() <-chan models.Odd {
	return a.updates
}

func (a *AxilisOfferFeed) processResponse(response *http.Response) {
	defer response.Body.Close()

	var odds []axilisOfferOdd
	err := json.NewDecoder(response.Body).Decode(&odds)
	if err != nil {
		log.Println("axilis offer feed, json decode", err)
		return
	}

	for _, odd := range odds {
		a.updates <- models.Odd{
			Id:          odd.Id,
			Name:        odd.Name,
			Match:       odd.Match,
			Coefficient: odd.Details.Price,
			Timestamp:   time.Now(),
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
