package messaging

import "time"

type Consumer interface {
	Receive(msg []byte) error
}

// EventReader run listen function inside goroutine to not block the main goroutine.
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
