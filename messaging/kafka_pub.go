package messaging

import (
	"context"

	kafka "github.com/segmentio/kafka-go"
)

func NewKafkaPublisher(config PublisherConfiguration) Publisher {
	writerConfig := kafka.WriterConfig{
		Brokers:  []string{config.BootstrapServer},
		Topic:    config.Topic,
		Balancer: &kafka.LeastBytes{},
		Async:    true,
	}

	return &kafkaWriter{
		writer: kafka.NewWriter(writerConfig),
	}
}

type kafkaWriter struct {
	writer *kafka.Writer
}

func (w *kafkaWriter) Publish(msg []byte) error {
	message := kafka.Message{
		Value: msg,
	}
	return w.writer.WriteMessages(context.Background(), message)
}

func (w *kafkaWriter) PublishWithContext(ctx context.Context, msg []byte) error {
	message := kafka.Message{
		Value: msg,
	}
	return w.writer.WriteMessages(ctx, message)
}
