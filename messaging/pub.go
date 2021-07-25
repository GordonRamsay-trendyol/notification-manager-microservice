package messaging

import "context"

// Publisher ...
type Publisher interface {
	Publish(message []byte) error
	PublishWithContext(ctx context.Context, msg []byte) error
}

type PublisherConfiguration struct {
	BootstrapServer string
	Topic           string
}
