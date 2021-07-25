package service

import (
	"encoding/json"
	"log"

	"github.com/GordonRamsay-trendyol/notification-manager-microservice/messaging"
	"github.com/GordonRamsay-trendyol/notification-manager-microservice/repository"
)

type NotificationType string

const (
	Email NotificationType = "EMAIL"
	Push  NotificationType = "PUSH"
	SMS   NotificationType = "SMS"
)

type NotificationService interface {
	Receive(msg []byte) error
}

func NewNotificationService() NotificationService {
	service := &notificationServiceImpl{
		smsPub:   messaging.NewPublisher(newPublisherConfig(SmsTopic)),
		emailPub: messaging.NewPublisher(newPublisherConfig(EmailTopic)),
		pushPub:  messaging.NewPublisher(newPublisherConfig(PushTopic)),
	}

	reader := messaging.NewReader(newEventReaderConfig(NotificationTopic), service)
	go reader.Listen()

	return service
}

type notificationServiceImpl struct {
	userRepository repository.UserRepository

	smsPub   messaging.Publisher
	emailPub messaging.Publisher
	pushPub  messaging.Publisher
}

type smsMsg struct {
	PhoneNumber string `json:"phoneNumber"`
	Message     string `json:"message"`
}

type emailMsg struct {
	EmailAddress string `json:"emailAddress"`
	Subject      string `json:"subject"`
	Content      string `json:"content"`
}

type pushMsg struct {
	PushID  string `json:"pushId"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type notificationMsg struct {
	To               int64  `json:"to"`
	Title            string `json:"title"`
	Content          string `json:"content"`
	NotificationType string `json:"type"`
}

func (publisher *notificationServiceImpl) sendSMS(phoneNumber, message string) error {
	sms := smsMsg{
		PhoneNumber: phoneNumber,
		Message:     message,
	}
	bytes, _ := json.Marshal(sms)
	log.Printf("Sms prepared, sms: %v\n", sms)
	return publisher.smsPub.Publish(bytes)
}

func (publisher *notificationServiceImpl) sendEmail(emailAddress, subject, content string) error {
	email := emailMsg{
		EmailAddress: emailAddress,
		Subject:      subject,
		Content:      content,
	}
	bytes, _ := json.Marshal(email)
	log.Printf("Email prepared, email: %v\n", email)
	return publisher.emailPub.Publish(bytes)
}

func (publisher *notificationServiceImpl) sendPushNotification(pushID, title, content string) error {
	msg := pushMsg{
		PushID:  pushID,
		Title:   title,
		Content: content,
	}
	bytes, _ := json.Marshal(msg)
	log.Printf("Push notification prepared, msg: %v\n", msg)
	return publisher.pushPub.Publish(bytes)
}

func (s *notificationServiceImpl) Receive(msg []byte) error {
	message := notificationMsg{}
	if err := json.Unmarshal(msg, &message); err != nil {
		return err
	}

	user, err := s.userRepository.FindById(message.To)

	if err != nil {
		log.Printf("Error occured, while fetching record from repository, err: %v\n, Failed message: %v\n", err, message)
		return err
	}

	if message.NotificationType == string(Email) && user.AllowEmail {
		s.sendEmail(user.Email, message.Title, message.Content)
	} else if message.NotificationType == string(SMS) && user.AllowSMS {
		s.sendSMS(user.PhoneNumber, message.Content)
	} else if message.NotificationType == string(Push) && user.AllowPushNotification {
		s.sendPushNotification(user.PushID, message.Title, message.Content)
	}

	return nil
}
