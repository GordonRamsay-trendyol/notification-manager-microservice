package service

import (
	"encoding/json"

	"github.com/GordonRamsay-Trendyol/notification-manager/messaging"
)

const (
	SmsTopic   = "sms_notification"
	EmailTopic = "email_notification"
	PushTopic  = "push_notification"
)

type NotificationService interface {
	SendSMS(phone, message string) error
	SendEmail(email, message string) error
	SendPushNotification(pushID, message string) error
}

func NewNotificationService() NotificationService {
	newConfig := func(topic string) messaging.PublisherConfiguration {
		return messaging.PublisherConfiguration{
			BootstrapServer: "localhost:9092",
			Topic:           topic,
		}
	}
	sms := newConfig(SmsTopic)
	email := newConfig(EmailTopic)
	push := newConfig(PushTopic)

	return &notificationServiceImpl{
		smsPub:   messaging.NewKafkaPublisher(sms),
		emailPub: messaging.NewKafkaPublisher(email),
		pushPub:  messaging.NewKafkaPublisher(push),
	}
}

type notificationServiceImpl struct {
	smsPub   messaging.Publisher
	emailPub messaging.Publisher
	pushPub  messaging.Publisher
}

func (publisher *notificationServiceImpl) SendSMS(phoneNumber, message string) error {
	sms := newSMS(phoneNumber, message)
	bytes, _ := json.Marshal(sms)
	return publisher.smsPub.Publish(bytes)
}

func (publisher *notificationServiceImpl) SendEmail(emailAddress, message string) error {
	email := newEmail(emailAddress, message)
	bytes, _ := json.Marshal(email)
	return publisher.emailPub.Publish(bytes)
}

func (publisher *notificationServiceImpl) SendPushNotification(pushID, message string) error {
	pushMsg := newPushMsg(pushID, message)
	bytes, _ := json.Marshal(pushMsg)
	return publisher.pushPub.Publish(bytes)
}

func newSMS(phoneNumber, message string) smsMessage {
	return smsMessage{
		phoneNumber: phoneNumber,
		message:     message,
	}
}

func newEmail(emailAddress, message string) emailMessage {
	return emailMessage{
		emailAddress: emailAddress,
		subject:      "CAMPAIGN MESSAGE",
		content:      message,
	}
}

func newPushMsg(pushID, message string) pushMessage {
	return pushMessage{
		pushID:      pushID,
		title:       "Campaign",
		description: message,
	}
}

type smsMessage struct {
	phoneNumber string
	message     string
}

type emailMessage struct {
	emailAddress string
	subject      string
	content      string
}

type pushMessage struct {
	pushID      string
	title       string
	description string
	// it can be many more...
}
