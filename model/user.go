package model

type User struct {
	ID int64

	Firstname string
	Lastname  string

	// Email will be used for email notifications
	Email string

	// PushID will be used for mobile push notifications
	PushID string

	// PhoneNumber will be used for sms notifications
	PhoneNumber string

	NotificationSettings NotificationSettings
}

type NotificationSettings struct {
	AllowPushNotification bool
	AllowSMS              bool
	AllowEmail            bool
}
