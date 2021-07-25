package service

import "github.com/GordonRamsay-trendyol/notification-manager-microservice/messaging"

const (
	DefaultBroker = "kafka:9092"

	NotificationTopic = "notification"
	SmsTopic          = "notification.sms"
	EmailTopic        = "notification.email"
	PushTopic         = "notification.push"

	UserUpdateTopic = "user.update"
)

func newEventReaderConfig(topic string) messaging.EventReaderConfig {
	return messaging.EventReaderConfig{
		Brokers:   []string{DefaultBroker},
		Topic:     topic,
		Partition: 3,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	}
}

func newPublisherConfig(topic string) messaging.PublisherConfiguration {
	return messaging.PublisherConfiguration{
		BootstrapServer: DefaultBroker,
		Topic:           topic,
	}
}
