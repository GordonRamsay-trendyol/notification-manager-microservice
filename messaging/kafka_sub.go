package messaging

import (
	"context"
	"log"

	kafka "github.com/segmentio/kafka-go"
)

type kafkaReader struct {
	reader *kafka.Reader
	sub    Consumer
}

func (c *kafkaReader) Listen() {

	for {
		m, err := c.reader.ReadMessage(context.Background())

		if err != nil {
			log.Printf("Listening failure, err: %v\n", err)
			break
		}

		if err := c.sub.Receive(m.Value); err != nil {
			log.Printf("callback function failed with err: %v", err)
		}

		log.Printf("message at offset %d: %s= %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := c.reader.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

func NewReader(config EventReaderConfig, sub Consumer) EventReader {
	readerConfig := kafka.ReaderConfig{
		Brokers:        config.Brokers,
		GroupID:        config.GroupID,
		Topic:          config.Topic,
		Partition:      config.Partition,
		MinBytes:       config.MinBytes,
		MaxBytes:       config.MaxBytes,
		CommitInterval: config.CommitInterval,
	}

	return &kafkaReader{
		reader: kafka.NewReader(readerConfig),
		sub:    sub,
	}
}
