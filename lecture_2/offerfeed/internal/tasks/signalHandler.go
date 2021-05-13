package tasks

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type SignalHandler struct{}

func NewSignalHandler() *SignalHandler {
	return &SignalHandler{}
}

func (s *SignalHandler) Start(ctx context.Context) error {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-signals:
		fmt.Println("caught signal")
	case <-ctx.Done():
	}

	return nil
}
