package queue

import (
	"context"
	"encoding/json"
	"os"

	"code-cadets-2021/lecture_2/offerfeed/internal/domain/models"

	"github.com/pkg/errors"
)

type OrderedQueue struct {
	queue  []models.Odd
	source chan models.Odd
}

func NewOrderedQueue() *OrderedQueue {
	return &OrderedQueue{
		source: make(chan models.Odd),
	}
}

func (o *OrderedQueue) Start(ctx context.Context) error {
	// on startup, load existing data from disk
	err := o.loadFromFile()
	if err != nil {
		return errors.Wrap(err, "ordered queue, load from file")
	}

	// cache all new odds
	for odd := range o.source {
		o.queue = append(o.queue, odd)
	}

	// on shutdown, persist everything to disk
	err = o.storeToFile()
	if err != nil {
		return errors.Wrap(err, "ordered queue, store to file")
	}

	return nil
}

func (o *OrderedQueue) GetSource() chan<- models.Odd {
	return o.source
}

func (o *OrderedQueue) loadFromFile() error {
	f, err := os.Open("queue.txt")
	if os.IsNotExist(err) {
		return nil

	} else if err != nil {
		return errors.Wrap(err, "load from file, open")
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&o.queue)
	if err != nil {
		return errors.Wrap(err, "load from file, decode")
	}

	return nil
}

func (o *OrderedQueue) storeToFile() error {
	f, err := os.Create("queue.txt")
	if err != nil {
		return errors.Wrap(err, "store to file, create")
	}
	defer f.Close()

	serializedQueue, err := json.MarshalIndent(o.queue, "", "    ")
	if err != nil {
		return errors.Wrap(err, "store to file, marshal")
	}

	_, err = f.Write(serializedQueue)
	if err != nil {
		return errors.Wrap(err, "store to file, write")
	}

	err = f.Sync()
	if err != nil {
		return errors.Wrap(err, "store to file, sync")
	}

	return nil
}
