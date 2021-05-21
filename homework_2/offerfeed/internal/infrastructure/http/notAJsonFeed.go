package http

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"code-cadets-2021/homework_2/offerfeed/internal/domain/models"
)

const notAJsonFeedURL = "http://18.193.121.232/axilis-feed-2"

type NotAJsonFeed struct {
	updates    chan models.Odd
	httpClient *http.Client
}

func NewNotAJsonFeed(httpClient *http.Client) *NotAJsonFeed {
	return &NotAJsonFeed{updates: make(chan models.Odd), httpClient: httpClient}
}

func (a *NotAJsonFeed) Start(ctx context.Context) error {
	defer close(a.updates)
	defer log.Printf("shutting down %s", a)

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(time.Second * 3):
			response, err := a.httpClient.Get(notAJsonFeedURL)
			if err != nil {
				log.Println("not a JSON feed, http get", err)
				continue
			}
			a.processResponse(ctx, response)
		}
	}
}

func (a *NotAJsonFeed) GetUpdates() chan models.Odd {
	return a.updates
}

func (a *NotAJsonFeed) String() string {
	return "not a JSON feed"
}

func (a *NotAJsonFeed) processResponse(ctx context.Context, response *http.Response) {
	defer response.Body.Close()

	bodyContent, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("not a JSON feed, io util read all", err)
		return
	}

	for _, line := range strings.Split(string(bodyContent), "\n") {
		parts := strings.Split(line, ",")

		coefficient, err := strconv.ParseFloat(parts[3], 64)
		if err != nil {
			log.Println("not a JSON feed, parse float", err)
		}

		odd := models.Odd{
			Id:          parts[0],
			Name:        parts[1],
			Match:       parts[2],
			Coefficient: coefficient,
			Timestamp:   time.Now(),
		}

		// IMPORTANT SELECT!!!
		// show an example
		select {
		case <-ctx.Done():
			return
		case a.updates <- odd:
			// do nothing
		}
	}
}
