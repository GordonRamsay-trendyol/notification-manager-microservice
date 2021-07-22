package messaging

import (
	"context"
	"log"
	"time"
)

type Configuration struct {
	BootstrapServers []string
	Topics           []string
	Partitions       int
}

// Start messaging listeners and publishers according to the given configuration.
func Start(config Configuration) {
	log.Println("Messaging services started..")
}

// Publisher ...
type Publisher interface {
	Publish(message []byte) error
	PublishWithContext(ctx context.Context, msg []byte) error
}

type PublisherConfiguration struct {
	BootstrapServer string
	Topic           string
}

type Consumer interface {
	Receive(msg []byte) error
}

type EventReader interface {
	Listen()
}

type EventReaderConfig struct {
	Brokers        []string
	GroupID        string
	Topic          string
	Partition      int
	MinBytes       int
	MaxBytes       int
	CommitInterval time.Duration
}
